package application

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/BenMeredithConsult/locagri-apps/app/adapters/presenters"
	"github.com/BenMeredithConsult/locagri-apps/ent"
	"github.com/BenMeredithConsult/locagri-apps/utils/bodyparser"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func OTP(size ...int) string {
	defaultSize := 6
	if size != nil {
		defaultSize = size[0]
	}
	chars := []byte("0123456789")
	b := make([]byte, defaultSize)
	_, _ = rand.Read(b)
	for i := 0; i < defaultSize; i++ {
		b[i] = chars[b[i]%byte(len(chars))]
	}
	return *(*string)(unsafe.Pointer(&b))
}
func RandomString(size int) string {
	chars := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, size)
	_, _ = rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = chars[b[i]%byte(len(chars))]
	}
	return *(*string)(unsafe.Pointer(&b))
}
func UsernameType(username, delimiter string) bool {
	if strings.Contains(username, "@") && delimiter == "email" {
		return true
	}
	phone, _ := regexp.Compile(`^0\d{9}$`)
	if phone.MatchString(username) && delimiter == "phone" {
		return true
	}
	return false
}
func Paginate(count int, results any) (*presenters.PaginationResponse, error) {
	return &presenters.PaginationResponse{
		Count: count,
		Data:  results,
	}, nil
}
func NoErrPaginate(count int, results any) *presenters.PaginationResponse {
	return &presenters.PaginationResponse{
		Count: count,
		Data:  results,
	}
}

func FormatSessionID(session any) int {
	if user, ok := session.(map[string]any); ok {
		userID, _ := strconv.Atoi(strconv.FormatFloat(user["id"].(float64), 'G', 'G', 64))
		return userID
	}
	return 0
}

func ConvertStructToMap(s any) map[string]interface{} {
	dataChan := make(chan map[string]interface{})
	go func(s any, dataChan chan map[string]interface{}) {
		v := reflect.ValueOf(s)
		values := make(map[string]interface{}, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			values[v.Type().Field(i).Name] = v.Field(i).Interface()
		}
		dataChan <- values
	}(s, dataChan)
	for {
		select {
		case values := <-dataChan:
			return values
		}
	}
}

func combinations(input []string, prefix []string, index int, result *[][]string) {
	if index == len(input) {
		if len(prefix) > 0 {
			*result = append(*result, append([]string{}, prefix...))
		}
		return
	}
	// Include the current element in the combination
	combinations(input, append(prefix, input[index]), index+1, result)

	// Exclude the current element from the combination
	combinations(input, prefix, index+1, result)
}

func FilterCombinations(input []string) [][]string {
	dataChan := make(chan [][]string)
	go func(dataChan chan [][]string) {
		var result [][]string
		combinations(input, []string{}, 0, &result)
		dataChan <- result
	}(dataChan)
	for {
		select {
		case result := <-dataChan:
			return result
		}
	}
}

func ParseRFC3339Datetime(rfc3339Datetime ...string) time.Time {
	if rfc3339Datetime == nil || rfc3339Datetime[0] == "" {
		return time.Now()
	}
	rfc3339Time, err := time.Parse(time.RFC3339, rfc3339Datetime[0])
	if err != nil {
		log.Panicln("Error parsing RFC3339 datetime:", err)
	}
	return rfc3339Time
}
func IsRFC3339Datetime(rfc3339Datetime string) bool {
	_, err := time.Parse(time.RFC3339, rfc3339Datetime)
	if err != nil {
		return false
	}
	return true
}
func IsTime(timeString string) bool {
	_, err := time.Parse(time.TimeOnly, timeString)
	if err != nil {
		return false
	}
	return true
}
func ParseRFC3339MYSQLDatetime(rfc3339Datetime string, format ...string) string {
	if rfc3339Datetime == "now" && format == nil {
		return time.Now().Format("2006-01-02")
	}

	if rfc3339Datetime == "now" && format != nil {
		return time.Now().Format(format[0])
	}
	if format != nil {
		rfc3339Time, err := time.Parse(time.RFC3339, rfc3339Datetime)
		if err != nil {
			log.Panicln("Error parsing RFC3339 datetime:", err)
		}
		return rfc3339Time.Format(format[0])
	}
	rfc3339Time, err := time.Parse(time.RFC3339, rfc3339Datetime)
	if err != nil {
		log.Panicln("Error parsing RFC3339 datetime:", err)
	}
	return rfc3339Time.Format("2006-01-02")
}

func CompareFilter(value any) bool {
	switch value.(type) {
	case bool:
		return value.(bool)
	case int:
		if value != 0 {
			return true
		}
	case string:
		if value != "" {
			return true
		}
	}
	return false
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
func RoundPercentage(val, base int, precision uint) float32 {
	percentage := float64(val) / float64(base) * 100
	ratio := math.Pow(10, float64(precision))
	return float32(math.Round(percentage*ratio) / ratio)
}
func BodyParser(c fiber.Ctx, v any) error {
	httpReq, err := adaptor.ConvertRequest(c, false)
	if err != nil {
		return err
	}
	return bodyparser.Parse(httpReq, v)
}

func HandleErrors(err error) (int, error) {
	if ent.IsNotFound(err) {
		return 400, fmt.Errorf("User not found")
	}
	switch err.Error() {
	case "User not found":
		return 400, fmt.Errorf("User not found")
	case "Invalid OTP":
		return 400, fmt.Errorf("Invalid OTP")
	case "OTP expired":
		return 400, fmt.Errorf("OTP expired")
	default:
		return 500, err
	}
}
