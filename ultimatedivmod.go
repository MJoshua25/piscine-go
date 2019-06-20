package piscine_go

func UltimateDivMod(a *int, b *int) {
                c := *a
               *a = *a / *b
                 
               *b = c % *b

                 
}