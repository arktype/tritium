package checker

import(
	tp "athena/proto"
	proto "goprotobuf.googlecode.com/hg/proto"
	. "strings"
)

func (result *CheckResult) CheckForSelectText(script *tp.ScriptObject) {
	iterate(script, func (ins *tp.Instruction) {
		if *ins.Type == tp.Instruction_FUNCTION_CALL {
			name := proto.GetString(ins.Value)
			if name == "$" || name == "select" {
				xpathIns := ins.Arguments[0] 
				xpath := proto.GetString(xpathIns.Value)
				if HasSuffix(xpath, "text()") {
					result.AddWarning(script, ins, "Shouldn't use '/text()' in a select() XPath.")
				}
			}
		}
		
	})
}

func (result *CheckResult) CheckForNotMisuse(script *tp.ScriptObject) {
	iterate(script, func (ins *tp.Instruction) {
		if *ins.Type == tp.Instruction_FUNCTION_CALL {
			name := proto.GetString(ins.Value)
			if name == "match" || name == "with" {
				if ins.Arguments != nil {
					for _, arg := range(ins.Arguments) {
						if *arg.Type == tp.Instruction_FUNCTION_CALL {
							if proto.GetString(arg.Value) == "not" {
								result.AddWarning(script, ins, "Possible misuse of not()– remember not is the opposite of with!")
							}
						}
					}
				}
				
			}
		}
		
	})
}