package tasktemplate

import "testing"

func TestTaskNoProject(t *testing.T) {

	expected := []Task{
		Task{"My sample 1", ""},
		Task{"My sample 2", ""},
		Task{"My sample 3", ""},
	}

	actuals := BuildTaskCommands("My sample #", "", 1, 3)

	for i, actual := range actuals {
		if actual != expected[i] {
			t.Errorf("%s is not %s", actual, expected[i])
		}
	}
}

func TestTaskWithProject(t *testing.T) {

	expected := []Task{
		Task{"My sample 1", "The Project"},
		Task{"My sample 2", "The Project"},
		Task{"My sample 3", "The Project"},
	}

	actuals := BuildTaskCommands("My sample #", "The Project", 1, 3)

	for i, actual := range actuals {
		if actual != expected[i] {
			t.Errorf("%s is not %s", actual, expected[i])
		}
	}
}

func TestTaskWeirdRange(t *testing.T) {

	expected := []Task{
		Task{"My sample 32", "The Project"},
		Task{"My sample 33", "The Project"},
		Task{"My sample 34", "The Project"},
		Task{"My sample 35", "The Project"},
		Task{"My sample 36", "The Project"},
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

func TestCommandArrayGen(t *testing.T) {
	expected := []string{"task", "add", "My sample 2", "project:\"The Project\""}
	actual := BuildTaskCommands("My sample #", "The Project", 2, 2)[0].GenerateTaskAddCMD()

	for i, item := range actual {
		if item != expected[i] {
			t.Errorf("%s is not %s", item, expected[i])
		}
	}
}
