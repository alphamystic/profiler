package entities

/*
  * This is a yara entity and is bound to change or evolve with time  say to one with imports
    This will vary as per with analysis and the mutates we get
*/
import (
  "github.com/alphamystic/profiler/libgo/utils"
)

// this is bound to change as time goes by
type YaraRule struct {
  YRId int
  Name string
  Meta []map[string]string
  Condition []string
  Actions []map[string]string
  utils.TimeStamps
}
