package strings

import "fmt"

func validatePlaceholders(format string, restrictToString bool) (int, error) {
	placeholderCount := 0

	for i := range len(format) {
		if i+1 == len(format) {
			break
		}

		if format[i] != '%' {
			continue
		}

		if i > 0 && format[i-1] == '%' {
			continue
		}

		if i < len(format) && format[i+1] == '%' {
			continue
		}

		if format[i+1] != 's' && restrictToString {
			return 0, fmt.Errorf("string format '%s' contains placeholders other than '%%s'", format)
		}

		placeholderCount += 1
	}

	return placeholderCount, nil
}

func SafeFormat(format string, values ...string) (string, error) {
	placeholderCount, err := validatePlaceholders(format, true)
	if err != nil {
		return "", err
	}

	if len(values) != placeholderCount {
		return "", fmt.Errorf("invalid number of values provided: expected %d but got %d", placeholderCount, len(values))
	}

	args := make([]any, len(values))
	for i, v := range values {
		args[i] = v
	}

	return fmt.Sprintf(format, args...), nil
}
