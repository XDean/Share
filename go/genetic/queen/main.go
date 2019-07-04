package queen

import "fmt"

func Main() {
	population := Population{
		Size:            100,
		Dim:             8,
		CrossoverFactor: 0.9,
		VariantFactor:   0.3,
		Target:          56,
	}.Random()
	result := Queens{}
outside:
	for {
		max := Queens{}
		for _, q := range population.Value {
			if q.TotalScore >= population.Target {
				result = q
				break outside
			}
			if q.TotalScore > max.TotalScore {
				max = q
			}
		}
		fmt.Printf("Gen %d, best score %d, value %v \n", population.Gen, max.TotalScore, max.Value)
		population = population.NextGen()
	}
	fmt.Println("Total Gen", population.Gen)
	fmt.Println("Answer", result.Value)
}
