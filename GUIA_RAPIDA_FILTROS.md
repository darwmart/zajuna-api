# 🚀 Guía Rápida de Uso - Filtro "Contiene"

## ⚡ Inicio Rápido (2 minutos)

### Paso 1: Seleccionar Tipo de Búsqueda
Cada campo tiene un selector dropdown con 5 opciones:
- **contiene** - Busca texto en cualquier parte del campo
- **no contiene** - Excluye registros que contengan el texto
- **es igual a** - Coincidencia exacta
- **comienza con** - Busca al inicio del campo
- **termina con** - Busca al final del campo

### Paso 2: Ingresar Valor
Escribe el texto a buscar en el campo de entrada

### Paso 3: Aplicar Filtro
Haz clic en el botón "Añadir filtro"

### Paso 4: Ver Resultados
Los resultados aparecen en la tabla y los filtros activos se muestran como pills arriba de la tabla

---

## 📋 Ejemplos Rápidos

### Buscar apellido "García"
```
Campo: Apellido(s)
Tipo: contiene
Valor: García
[Añadir filtro]
```

### Excluir emails de Gmail
```
Campo: Dirección de correo
Tipo: no contiene
Valor: gmail
[Añadir filtro]
```

### Buscar username exacto
```
Campo: Nombre de usuario
Tipo: es igual a
Valor: admin
[Añadir filtro]
```

### Buscar profesores (username comienza con "prof")
```
Campo: Nombre de usuario
Tipo: comienza con
Valor: prof
[Añadir filtro]
```

### Buscar emails institucionales (termina con dominio)
```
Campo: Dirección de correo
Tipo: termina con
Valor: @universidad.edu
[Añadir filtro]
```

---

## 🔗 Combinar Múltiples Filtros

Los filtros se combinan con lógica AND (deben cumplirse TODOS)

**Ejemplo:** Profesores activos con email institucional
```
1. Username | comienza con | prof_
2. Email | termina con | @universidad.edu
3. Username | no contiene | _inactive
```

---

## 🎯 Tips Rápidos

✅ **Búsqueda flexible**: Usa "contiene" para búsquedas generales  
✅ **Filtrar datos**: Usa "no contiene" para excluir  
✅ **Búsqueda precisa**: Usa "es igual a" para exactitud  
✅ **Patrones**: Usa "comienza con" y "termina con" para patrones  

---

## 🗑️ Eliminar Filtros

### Eliminar un filtro específico
Haz clic en la **X** de la pill del filtro

### Eliminar todos los filtros
Haz clic en "Limpiar todos los filtros"

---

## ❓ FAQ Rápido

**P: ¿Es case-sensitive?**  
R: No, las búsquedas ignoran mayúsculas/minúsculas

**P: ¿Puedo buscar con espacios?**  
R: Sí, los espacios son parte de la búsqueda

**P: ¿Cuántos filtros puedo aplicar?**  
R: Ilimitados, pero se recomienda máximo 5

**P: ¿Los filtros se guardan?**  
R: No, se resetean al recargar la página

---

## 🎨 Interfaz Visual

```
┌─────────────────────────────────────────────────────────┐
│ Nuevo filtro                                            │
│                                                         │
│ Nombre completo [contiene ▼] [Buscar usuario......]    │
│ Mostrar más...                                          │
│                      [Añadir filtro]                    │
└─────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────┐
│ Filtros activos:                                        │
│ [Apellido(s) contiene "García" ×]                       │
│ [Email termina con "@gmail.com" ×]                      │
│ [Limpiar todos los filtros]                             │
└─────────────────────────────────────────────────────────┘
```

---

## ⌨️ Atajos de Teclado

- `Tab` - Navegar entre campos
- `Enter` - Aplicar filtro (desde input)
- `Esc` - Cerrar selector (si está abierto)

---

## 📱 Uso en Móvil

✅ Totalmente responsive  
✅ Selectores táctiles optimizados  
✅ Pills con área de toque amplia  

---

## ✅ Checklist de Uso

- [ ] Seleccionar campo a filtrar
- [ ] Elegir tipo de coincidencia
- [ ] Ingresar valor de búsqueda
- [ ] Hacer clic en "Añadir filtro"
- [ ] Verificar pill de filtro activo
- [ ] Revisar resultados en tabla
- [ ] (Opcional) Agregar más filtros
- [ ] (Opcional) Eliminar filtros con X

---

## 🆘 Solución Rápida de Problemas

**No hay resultados:**
- Verifica que el tipo de coincidencia es correcto
- Intenta con "contiene" en vez de "es igual a"
- Revisa que no haya espacios extras

**Muchos resultados:**
- Agrega más filtros para refinar
- Usa tipos más específicos ("es igual a", "comienza con")

**Filtro no se aplica:**
- Asegúrate de hacer clic en "Añadir filtro"
- Verifica que el valor no esté vacío

---

## 🎓 Para Usuarios de Moodle

**¡Funciona exactamente igual que en Moodle!**

Si ya usas filtros en Moodle, no necesitas aprender nada nuevo.  
La interfaz y el comportamiento son idénticos.

---

**¿Necesitas más ayuda?**  
📖 Consulta: `EJEMPLOS_FILTRO_CONTIENE.md`  
🔧 Documentación técnica: `FILTRO_CONTIENE_DOCUMENTACION.md`

---

**¡Happy Filtering! 🎉**
