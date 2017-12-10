package tasktemplate

import "testing"

func TestTaskNoProject(t *testing.T) {

	expected := []string{
		"task add \"My sample 1\" project:\"\"",
		"task add \"My sample 2\" project:\"\"",
		"task add \"My sample 3\" project:\"\"",
	}

	actuals := BuildTaskCommands("My sample #", "", 1, 3)

	for i, actual := range actuals {
		if actual != expected[i] {
			t.Errorf("%s is not %s", actual, expected[i])
		}
	}
}

func TestTaskWithProject(t *testing.T) {

	expected := []string{
		"task add \"My sample 1\" project:\"The Project\"",
		"task add \"My sample 2\" project:\"The Project\"",
		"task add \"My sample 3\" project:\"The Project\"",
	}

	actuals := BuildTaskCommands("My sample #", "The Project", 1, 3)

	for i, actual := range actuals {
		if actual != expected[i] {
			t.Errorf("%s is not %s", actual, expected[i])
		}
	}
}

func TestTaskWeirdRange(t *testing.T) {

	expected := []string{
		"task add \"My sample 32\" project:\"The Project\"",
		"task add \"My sample 33\" project:\"The Project\"",
		"task add \"My sample 34\" project:\"The Project\"",
		"task add \"My sample 35\" project:\"The Project\"",
		"task add \"My sample 36\" project:\"The Project\"",
	}

	actuals := BuildTaskCommands("My sample #", "The Project", 32, 36)

	for i, actual := range actuals {
		if actual != expected[i] {
			t.Errorf("%s is not %s", actual, expected[i])
		}
	}
}

func TestTaskBadRange(t *testing.T) {

	actuals := BuildTaskCommands("My sample #", "The Project", 36, 32)

	if len(actuals) != 0 {
		t.Errorf("%s is not empty", actuals)
	}
}
