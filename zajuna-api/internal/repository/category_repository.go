package repository

import (
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
