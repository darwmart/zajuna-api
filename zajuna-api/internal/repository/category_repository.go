package repository

import (
	"fmt"
	"zajunaApi/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category

	if err := r.db.Table("mdl_course_categories").
		Order("sortorder").
		Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// MoveCategory mueve una categoría antes de otra categoría especificada
// Soporta cambio de orden (reordenamiento) y cambio de padre (parentid)
func (r *CategoryRepository) MoveCategory(id uint, beforeid uint, parentid *uint) error {
	// 1. Verificar que la categoría a mover existe
	var categoryToMove models.Category
	if err := r.db.Table("mdl_course_categories").
		Where("id = ?", id).
		First(&categoryToMove).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}

	// Determinar el nuevo padre: si parentid es nil, mantener el padre actual
	var newParent uint
	if parentid != nil {
		newParent = *parentid

		// 2. Si cambia de padre, verificar que el nuevo padre existe
		if newParent != uint(categoryToMove.Parent) && newParent != 0 {
			var parentCategory models.Category
			if err := r.db.Table("mdl_course_categories").
				Where("id = ?", newParent).
				First(&parentCategory).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return gorm.ErrRecordNotFound
				}
				return err
			}
		}
	} else {
		// Mantener el padre actual
		newParent = uint(categoryToMove.Parent)
	}

	// 3. Si beforeid != 0, verificar que existe y tiene el mismo nuevo parent
	if beforeid != 0 {
		var targetCategory models.Category
		if err := r.db.Table("mdl_course_categories").
			Where("id = ?", beforeid).
			First(&targetCategory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return gorm.ErrRecordNotFound
			}
			return err
		}

		// Verificar que la categoría objetivo tiene el mismo nuevo padre
		if uint(targetCategory.Parent) != newParent {
			return gorm.ErrInvalidValue
		}
	}

	// 4. Obtener todas las categorías del nuevo padre, ordenadas por sortorder
	var siblings []models.Category
	if err := r.db.Table("mdl_course_categories").
		Where("parent = ?", newParent).
		Order("sortorder ASC, id ASC").
		Find(&siblings).Error; err != nil {
		return err
	}

	// 5. Construir la nueva lista de categorías reordenadas
	var reordered []models.Category

	if beforeid == 0 {
		// Mover al final
		for _, cat := range siblings {
			if cat.ID != id {
				reordered = append(reordered, cat)
			}
		}
		reordered = append(reordered, categoryToMove)
	} else {
		// Encontrar la posición actual de la categoría a mover
		var currentIndex int
		for i, cat := range siblings {
			if cat.ID == id {
				currentIndex = i
				break
			}
		}

		// Encontrar la posición del beforeid
		var beforeIndex int
		for i, cat := range siblings {
			if cat.ID == beforeid {
				beforeIndex = i
				break
			}
		}

		// Determinar la categoría adyacente con la que intercambiar
		var adjacentIndex int
		if beforeIndex < currentIndex {
			// Mover hacia arriba: intercambiar con la anterior (currentIndex - 1)
			adjacentIndex = currentIndex - 1
		} else {
			// Mover hacia abajo: intercambiar con la siguiente (currentIndex + 1)
			adjacentIndex = currentIndex + 1
		}

		// Construir el nuevo array intercambiando posiciones
		for i, cat := range siblings {
			if i == currentIndex {
				// En la posición de la categoría a mover, poner la adyacente
				reordered = append(reordered, siblings[adjacentIndex])
			} else if i == adjacentIndex {
				// En la posición de la adyacente, poner la categoría a mover
				reordered = append(reordered, categoryToMove)
			} else {
				// Las demás se mantienen igual
				reordered = append(reordered, cat)
			}
		}
	}
	// 6. Si cambió de padre, actualizar el parentid
	if parentid != nil && uint(categoryToMove.Parent) != newParent {
		if err := r.db.Table("mdl_course_categories").
			Where("id = ?", id).
			Update("parent", newParent).Error; err != nil {
			return err
		}
	}

	// 7. Actualizar sortorder secuencialmente para todas las categorías del nuevo padre
	// y actualizar el sortOrder en el slice para usarlo después
	for i := range reordered {
		newSortOrder := (i + 1) * 10000
		if err := r.db.Table("mdl_course_categories").
			Where("id = ?", reordered[i].ID).
			Update("sortorder", newSortOrder).Error; err != nil {
			return err
		}
		// Actualizar el sortOrder en el slice para reflejar el nuevo valor
		reordered[i].SortOrder = newSortOrder
	}

	// 8. Actualizar sortorder de los cursos dentro de las categorías afectadas
	// Ahora 'reordered' tiene los sortOrder actualizados
	if err := r.UpdateCoursesSortOrderForCategories(reordered); err != nil {
		return err
	}

	return nil
}

// UpdateCoursesSortOrderForCategories actualiza el sortorder de los cursos dentro de las categorías dadas
// Los cursos heredan el sortorder de su categoría padre para mantener el orden global
func (r *CategoryRepository) UpdateCoursesSortOrderForCategories(categories []models.Category) error {
	for _, category := range categories {
		// Obtener todos los cursos de esta categoría ordenados por su sortorder actual
		var courses []models.Course
		if err := r.db.Table("mdl_course").
			Where("category = ?", category.ID).
			Order("sortorder ASC, id ASC").
			Find(&courses).Error; err != nil {
			return err
		}

		// Actualizar el sortorder de cada curso basándose en el sortorder de la categoría
		// Cada curso tiene sortorder = sortorder_categoria + posición (autoincremental por unidad)
		for i, course := range courses {
			// Nuevo sortorder = sortorder_categoria + (índice + 1)
			// Ejemplo: categoria=60000 -> curso1=60001, curso2=60002, curso3=60003
			newCourseSortOrder := category.SortOrder + (i + 1)

			if err := r.db.Table("mdl_course").
				Where("id = ?", course.ID).
				Update("sortorder", newCourseSortOrder).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// CreateCategories crea una o más categorías siguiendo las reglas de Moodle 4.3
func (r *CategoryRepository) CreateCategories(categories []models.Category) ([]models.Category, error) {
	var createdCategories []models.Category

	for _, category := range categories {
		// 1. Validar que el padre existe (si parent > 0)
		if category.Parent > 0 {
			var parentCategory models.Category
			if err := r.db.Table("mdl_course_categories").
				Where("id = ?", category.Parent).
				First(&parentCategory).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return nil, gorm.ErrRecordNotFound
				}
				return nil, err
			}
			// Calcular depth y path basados en el padre
			category.Depth = parentCategory.Depth + 1
			category.Path = parentCategory.Path + "/" + "0" // Será actualizado después de la creación
		} else {
			// Categoría de nivel superior
			category.Depth = 1
			category.Path = "/0" // Será actualizado después de la creación
		}

		// 2. Establecer valores por defecto
		if category.Visible == 0 && category.Visible != 1 {
			category.Visible = 1 // Por defecto visible
		}
		if category.DescriptionFormat == 0 {
			category.DescriptionFormat = 1 // HTML por defecto
		}

		// 3. Calcular el sortorder: obtener el máximo sortorder de las categorías hermanas + 10000
		var maxSortOrder int
		err := r.db.Table("mdl_course_categories").
			Where("parent = ?", category.Parent).
			Select("COALESCE(MAX(sortorder), 0)").
			Scan(&maxSortOrder).Error
		if err != nil {
			return nil, err
		}
		category.SortOrder = maxSortOrder + 10000

		// 4. Crear la categoría en la base de datos
		if err := r.db.Table("mdl_course_categories").Create(&category).Error; err != nil {
			return nil, err
		}

		// 5. Actualizar el path con el ID real
		if category.Parent > 0 {
			// Obtener el path del padre y concatenar el ID de la nueva categoría
			var parentPath string
			r.db.Table("mdl_course_categories").
				Where("id = ?", category.Parent).
				Select("path").
				Scan(&parentPath)
			category.Path = fmt.Sprintf("%s/%d", parentPath, category.ID)
		} else {
			category.Path = fmt.Sprintf("/%d", category.ID)
		}

		// 6. Actualizar el path en la base de datos
		if err := r.db.Table("mdl_course_categories").
			Where("id = ?", category.ID).
			Update("path", category.Path).Error; err != nil {
			return nil, err
		}

		createdCategories = append(createdCategories, category)
	}

	return createdCategories, nil
}
