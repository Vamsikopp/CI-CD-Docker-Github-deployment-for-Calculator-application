




// Even return true if the number of element in the Cell is even.
func (c *Cell) Even() bool {
	return c.even
}

func (c *Cell) validate() bool {
	if c.err != nil {
		return false
	}
	return true
}

// Total returns the total value of values in Cell.
func (c *Cell) Total() *Cell {
	if c.calculatedTotal || c.err != nil {
		return c
	}
	c.Register.Total = 0
	for _, v := range c.data {
		c.Register.Total += v
	}
	c.calculatedTotal = true
	return c
}

// Mean returns the mean (average) value of values in Cell.
func (c *Cell) Mean() *Cell {
	if c.calculatedMean || c.err != nil {
		return c
	}
	if !c.calculatedTotal {
		c.Total()
		if c.err != nil {
			return c
		}
	}
	c.Register.Mean = c.Register.Total / float64(c.length)
	c.calculatedMean = true
	return c
}

// Variance calculates variance of the values in Cell.
func (c *Cell) Variance() *Cell {
	if c.calculatedVariance || c.err != nil {
		return c
	}
	if !c.calculatedMean {
		c.Mean()
		if c.err != nil {
			return c
		}
	}
	for _, i := range c.data {
		r := c.Register.Mean - i
		c.Register.Variance += r * r
	}
	c.Register.Variance /= float64(c.length)
	c.calculatedVariance = true
	return c
}

// Range calculates range of the values in Cell.
func (c *Cell) Range() *Cell {
	if c.calculatedRange || c.err != nil {
		return c
	}

	if !c.calculatedMax {
		c.Max()
		if c.err != nil {
			return c
		}
	}

	if !c.calculatedMin {
		c.Min()
		if c.err != nil {
			return c
		}
	}

	c.Register.Range = c.Register.MaxValue - c.Register.MinValue
	c.calculatedRange = true
	return c
}

// Max calculates the biggest value in Cell.
func (c *Cell) Max() *Cell {
	if c.calculatedMax || c.err != nil {
		return c
	}
	c.Register.MaxValue = c.sortedData[c.length-1]
	c.calculatedMax = true
	return c
}

// MaxWithIndices calculates the biggest value and associated indices in Cell.
func (c *Cell) MaxWithIndices() *Cell {
	if c.calculatedMaxWithIndices || c.err != nil {
		return c
	}

	if !c.calculatedMax {
		c.Max()
		if c.err != nil {
			return c
		}
	}
	c.Register.MaxIndices = []int{}
	if c.length == 1 {
		c.Register.MaxIndices = append(c.Register.MaxIndices, 0)
		c.calculatedMaxWithIndices = true
		return c
	}
	for i, v := range c.data {
		if v == c.Register.MaxValue {
			c.Register.MaxIndices = append(c.Register.MaxIndices, i)
		}
	}
	c.calculatedMaxWithIndices = true
	return c
}

// Min calculates the smallest value in Cell.
func (c *Cell) Min() *Cell {
	if c.calculatedMin || c.err != nil {
		return c
	}
	c.Register.MinValue = c.sortedData[0]
	c.calculatedMax = true
	return c
}

// MinWithIndices calculates the smallest value and associated indices in Cell.
func (c *Cell) MinWithIndices() *Cell {
	if c.calculatedMinWithIndices || c.err != nil {
		return c
	}

	if !c.calculatedMin {
		c.Min()
		if c.err != nil {
			return c
		}
	}
	c.Register.MinIndices = []int{}
	if c.length == 1 {
		c.Register.MinIndices = append(c.Register.MinIndices, 0)
		c.calculatedMinWithIndices = true
		return c
	}
	for i, v := range c.data {
		if v == c.Register.MinValue {
			c.Register.MinIndices = append(c.Register.MinIndices, i)
		}
	}
	c.calculatedMinWithIndices = true
	return c
}

// Modes calculates the values appearing most often in Cell.
func (c *Cell) Modes() *Cell {
	if c.calculatedModes || c.err != nil {
		return c
	}
	occurences := make(map[float64]int)
	c.Register.Modes = []float64{}
	for _, i := range c.data {
		occurences[i]++
		if occurences[i] > c.Register.ModeRepeatCount {
			c.Register.ModeRepeatCount = occurences[i]
		}
	}
	if len(occurences) == c.length {
		c.Register.ModeRepeatCount = 0
		c.calculatedModes = true
		return c
	}
	for i, v := range occurences {
		if v == c.Register.ModeRepeatCount {
			c.Register.Modes = append(c.Register.Modes, i)
		}
	}
	c.calculatedModes = true
	return c
}

// Median calculates median of the values in Cell.
func (c *Cell) Median(sorted bool) *Cell {
	if (!sorted && c.calculatedMedian) || (sorted && c.calculatedSortedMedian) || c.err != nil {
		return c
	}
	if c.length == 1 {
		c.Register.SortedMedian = c.data[0]
		c.Register.Median = c.data[0]
		c.calculatedSortedMedian = true
		c.calculatedMedian = true
		return c
	}

	if sorted {
		if !c.even {
			c.Register.SortedMedian = c.sortedData[c.middleIndex]
		} else {
			c.Register.SortedMedian = (c.sortedData[c.middleIndex] + c.sortedData[c.middleIndex-1]) / 2
		}
		c.calculatedSortedMedian = true
	} else {
		if !c.even {
			c.Register.Median = c.data[c.middleIndex]
		} else {
			c.Register.Median = (c.data[c.middleIndex] + c.data[c.middleIndex-1]) / 2
		}
		c.calculatedMedian = true
	}
	return c
}

// StandardDeviation calculates standard deviation of the values in Cell.
func (c *Cell) StandardDeviation() *Cell {
	if c.calculatedStandardDeviation || c.err != nil {
		return c
	}

	if !c.calculatedVariance {
		c.Variance()
		if c.err != nil {
			return c
		}
	}

	c.Register.StandardDeviation = math.Sqrt(c.Register.Variance)
	c.calculatedStandardDeviation = true
	return c
}

// RunAll performs all available culculations in Cell.
func (c *Cell) RunAll() *Cell {
	if valid := c.validate(); !valid {
		return c



// NewUint32 returns an instance of Cell from uint32 array.
func NewUint32(data []uint32) *Cell {
	arr := []float64{}
	for _, i := range data {
		arr = append(arr, float64(i))
	}
	return New(arr)
}

}
