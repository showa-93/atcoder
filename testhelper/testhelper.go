package testhelper

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type RandomTestFunc func() string

type RandomTestBuilder struct {
	buildFuncMap map[string]RandomTestFunc
	buildOrders  []string
	data         map[string]interface{}
}

func NewRandomTestBuilder() *RandomTestBuilder {
	return &RandomTestBuilder{
		buildFuncMap: make(map[string]RandomTestFunc),
		buildOrders:  make([]string, 0),
		data:         make(map[string]interface{}),
	}
}

// AddInt sets the rondom value for the test condition. The range of values is [from, to).
func (b *RandomTestBuilder) AddIntKey(key string, from, to interface{}) {
	rand.Seed(time.Now().UnixMicro())
	b.buildFuncMap[key] = func() string {
		from2, to2 := b.convertInt(from), b.convertInt(to)
		if from2 >= to2 {
			panic(`arguments must set "from < to"`)
		}
		return strconv.Itoa(rand.Intn(to2-from2+1) + from2)
	}
}

func (b *RandomTestBuilder) AddBuildOrder(count interface{}, keys []string) {
	builderKey := strings.Join(keys, " ")
	b.buildFuncMap[builderKey] = func() string {
		n := b.convertInt(count)
		lines := make([]string, 0, n)
		for i := 0; i < n; i++ {
			values := make([]string, len(keys))
			for j, key := range keys {
				buildFunc, ok := b.buildFuncMap[key]
				if !ok {
					panic(fmt.Sprintf("%s does not exist. key=%s", key, builderKey))
				}
				values[j] = buildFunc()
				b.data[key] = values[j]
			}
			lines = append(lines, strings.Join(values, " "))
		}

		return strings.Join(lines, "\n")
	}

	b.buildOrders = append(b.buildOrders, builderKey)
}

func (b *RandomTestBuilder) Build() string {
	var sb strings.Builder
	for _, order := range b.buildOrders {
		buildFunc := b.buildFuncMap[order]
		sb.WriteString(buildFunc())
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (b *RandomTestBuilder) convertInt(v interface{}) int {
	switch x := v.(type) {
	case string:
		value, ok := b.data[x]
		if !ok {
			panic(fmt.Sprintf("%s does not exist.", x))
		}
		return convertInt(value)
	case int:
		return x
	default:
		panic(fmt.Sprintf("unsuppoted type. type=%s", x))
	}
}

func convertInt(v interface{}) int {
	switch x := v.(type) {
	case string:
		y, _ := strconv.Atoi(x)
		return y
	case int:
		return x
	default:
		panic(fmt.Sprintf("unsuppoted type. type=%s", x))
	}
}
