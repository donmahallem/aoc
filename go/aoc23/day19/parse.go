package day19

import (
	"bufio"
	_ "embed"
	"io"
	"strconv"
	"strings"
)

//go:embed sample.txt
var sample string

const actionReject = "R"
const actionAccept = "A"

type workflowRule interface {
	Evaluate(r rating) (string, bool)
}

type workflowRuleGreater struct {
	letter byte
	value  int
	target string
}

func (wr *workflowRuleGreater) Evaluate(r rating) (string, bool) {
	value, ok := r.valueFor(wr.letter)
	if !ok {
		return "", false
	}
	if value > wr.value {
		return wr.target, true
	}
	return "", false
}

type workflowRuleLess struct {
	letter byte
	value  int
	target string
}

func (wr *workflowRuleLess) Evaluate(r rating) (string, bool) {
	value, ok := r.valueFor(wr.letter)
	if !ok {
		return "", false
	}
	if value < wr.value {
		return wr.target, true
	}
	return "", false
}

type workflowRuleDirect struct {
	target string
}

func (wr *workflowRuleDirect) Evaluate(r rating) (string, bool) {
	return wr.target, true
}

// holding workflow rules
type workflow struct {
	rules []workflowRule
}

func (wf workflow) Next(r rating) (string, bool) {
	for _, rule := range wf.rules {
		if target, ok := rule.Evaluate(r); ok {
			return target, true
		}
	}
	return "", false
}

type rating struct {
	X int
	M int
	A int
	S int
}

func (r rating) valueFor(letter byte) (int, bool) {
	switch letter {
	case 'x':
		return r.X, true
	case 'm':
		return r.M, true
	case 'a':
		return r.A, true
	case 's':
		return r.S, true
	default:
		return 0, false
	}
}

type parsedData struct {
	Workflows map[string]workflow
	Ratings   []rating
}

func parseInput(r io.Reader) parsedData {
	scanner := bufio.NewScanner(r)
	data := parsedData{
		Workflows: make(map[string]workflow),
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		openIdx := strings.IndexByte(line, '{')
		closeIdx := strings.LastIndexByte(line, '}')
		if openIdx == -1 || closeIdx == -1 || closeIdx <= openIdx {
			continue
		}

		name := line[:openIdx]
		rulesPart := line[openIdx+1 : closeIdx]
		rawRules := strings.Split(rulesPart, ",")
		rules := make([]workflowRule, 0, len(rawRules))

		for _, rawRule := range rawRules {
			rawRule = strings.TrimSpace(rawRule)
			if rawRule == "" {
				continue
			}

			if sepIdx := strings.IndexByte(rawRule, ':'); sepIdx != -1 {
				cond := rawRule[:sepIdx]
				target := rawRule[sepIdx+1:]
				if len(cond) < 3 {
					continue
				}

				letter := cond[0]
				op := cond[1]
				valStr := cond[2:]
				val, err := strconv.Atoi(valStr)
				if err != nil {
					continue
				}

				switch op {
				case '<':
					rules = append(rules, &workflowRuleLess{
						letter: letter,
						value:  val,
						target: target,
					})
				case '>':
					rules = append(rules, &workflowRuleGreater{
						letter: letter,
						value:  val,
						target: target,
					})
				}
			} else {
				rules = append(rules, &workflowRuleDirect{target: rawRule})
			}
		}

		data.Workflows[name] = workflow{rules: rules}
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		rating, ok := parseRating(line)
		if ok {
			data.Ratings = append(data.Ratings, *rating)
		}
	}

	return data
}

func parseRating(line string) (*rating, bool) {
	if len(line) < 2 || line[0] != '{' || line[len(line)-1] != '}' {
		return nil, false
	}

	r := rating{}
	inner := line[1 : len(line)-1]
	parts := strings.Split(inner, ",")
	if len(parts) != 4 {
		return nil, false
	}

	for _, part := range parts {
		kv := strings.SplitN(part, "=", 2)
		if len(kv) != 2 {
			return nil, false
		}

		key := kv[0]
		value, err := strconv.Atoi(kv[1])
		if err != nil {
			return nil, false
		}

		switch key {
		case "x":
			r.X = value
		case "m":
			r.M = value
		case "a":
			r.A = value
		case "s":
			r.S = value
		default:
			return nil, false
		}
	}

	return &r, true
}
