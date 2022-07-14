package main
	
func TestSoma(t * testing.T){
	total := Soma(15,15)
	if total!=30{
		t.Fatalf("Resultado da Soma é  invalido: Resultado %v. Esperado %v", total, 30)
	}
}
