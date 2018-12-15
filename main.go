//Robert Simons
//CptS 484
//Individual Project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//This application takes in a specified population number (entities) with a corresponding specified number of genes.
//and runs them repetitavely through a genetic algorithm until they are all 1's.
func main() {

	generationCount := 0
	valid := true
	tempEntity := make([]entity, 10)

	//create individuals in the population
	for j := range tempEntity {
		rand.Seed(time.Now().UnixNano())

		tempGene := make([]gene, 5)

		//create genes for each individual in the population
		for i := range tempGene {
			tempGene[i].binary = random(1000, 2000) % 2
		}

		//individual attributes. Can be changed. Currently each individual has 5 genes.
		e := entity{
			fitness:    0,
			geneLength: 5,
			genes:      tempGene,
		}

		tempEntity[j] = e
	}

	//population attributes. Can be changed. Currently there are 10 individuals in the population
	p := population{
		size:     10,
		mostFit:  0,
		entities: tempEntity,
	}

	/*
		Start
		Generate initial population
		Compute fitness
		repeat:
			   selection
			   crossover
			   mutation
		Until poulation has converged (all genes are the binary 1)
		Stop
	*/
	for valid {
		generationCount++
		fittest, secondFittest, leastFitIndex := p.getFitCouple()

		//All genes are the binary 1
		if fittest.calcFitness() >= 5 {
			valid = false
			goto Exit
		}

		crossoverPoint := random(0, 5)

		//Swap genes between the fittest and second fittest individual
		for i := 0; i < crossoverPoint; i++ {
			tempGene := fittest.genes[i].binary
			fittest.genes[i].binary = secondFittest.genes[i].binary
			secondFittest.genes[i].binary = tempGene
		}

		//Change a gene at random with the mutation currently set to if the RandomNum % 27 < 3
		mutationRandom := random(0, 2000)
		if mutationRandom%27 < 3 {
			mutationPoint := random(0, 5)
			if secondFittest.genes[mutationPoint].binary == 0 {
				secondFittest.genes[mutationPoint].binary = 1
			} else {
				secondFittest.genes[mutationPoint].binary = 0
			}
		}

		//getting fittest offspring and replacing the least fit with it
		if fittest.calcFitness() > secondFittest.calcFitness() {
			fittestOffspring := fittest
			p.entities[leastFitIndex] = fittestOffspring
		} else {
			fittestOffspring := secondFittest
			p.entities[leastFitIndex] = fittestOffspring
		}
	}

Exit: //Algorithm has Completed its iterations
	fmt.Println("Solution found in generation", generationCount)
}

//generate random numbers between values
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

//default attribute of gene
type gene struct {
	binary int
}

//default attribute of entity (individual)
type entity struct {
	fitness    int
	geneLength int
	genes      []gene
}

//sum of all the individuals genes is their fitness. When they are all 1 they are 'ideal'
func (e entity) calcFitness() int {

	for _, gene := range e.genes {
		e.fitness += gene.binary
	}
	return e.fitness
}

//default attribute of population
type population struct {
	size     int
	mostFit  int
	entities []entity
}

//finding the fittest individual, second fittest individual, and least fit individual in the population
func (p population) getFitCouple() (entity, entity, int) {
	mostFit1, mostFit2, itr := 0, 0, 0
	leastFit := 100000
	leastFitIndex := 0

	for _, entity := range p.entities {
		if leastFit >= entity.calcFitness() {
			leastFit = entity.calcFitness()
			leastFitIndex = itr
		}
		if entity.calcFitness() > p.entities[mostFit1].calcFitness() {
			mostFit2 = mostFit1
			mostFit1 = itr
		} else if entity.calcFitness() > p.entities[mostFit2].calcFitness() {
			mostFit2 = itr
		}
		itr++
	}
	return p.entities[mostFit1], p.entities[mostFit2], leastFitIndex
}
