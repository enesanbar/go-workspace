package gonum_basics

import (
	"fmt"

	"github.com/gonum/blas/blas64"
	"github.com/gonum/floats"
	"github.com/gonum/matrix/mat64"
)

func Vectors() {

	// Create a new vector value.
	myvector := mat64.NewVector(2, []float64{11.0, 5.2})
	fmt.Println(myvector)

	// Initialize a couple of "vectors" represented as slices.
	vectorSliceA := []float64{11.0, 5.2, -1.3}
	vectorSliceB := []float64{-7.2, 4.2, 5.1}

	// Compute the dot product of A and B
	// (https://en.wikipedia.org/wiki/Dot_product).
	dotProduct := floats.Dot(vectorSliceA, vectorSliceB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)

	// Scale each element of A by 1.5.
	floats.Scale(1.5, vectorSliceA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorSliceA)

	// Compute the norm/length of B.
	normB := floats.Norm(vectorSliceB, 2)
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)

	// Initialize a couple of "vectors" represented as slices.
	vectorA := mat64.NewVector(3, []float64{11.0, 5.2, -1.3})
	vectorB := mat64.NewVector(3, []float64{-7.2, 4.2, 5.1})

	// Compute the dot product of A and B
	// (https://en.wikipedia.org/wiki/Dot_product).
	dotProduct = mat64.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)

	// Scale each element of A by 1.5.
	vectorA.ScaleVec(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)

	// Compute the norm/length of B.
	normB = blas64.Nrm2(3, vectorB.RawVector())
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
}
