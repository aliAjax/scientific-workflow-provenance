import json
from pathlib import Path

root = Path(__file__).resolve().parents[3]
date = "2026-08-21"
plan = json.loads((root / "_shared" / "bug_plan.json").read_text(encoding="utf-8"))
for idx, bug in enumerate(plan["bugs"], 1):
    rec = f"{idx:03d}"
    path = root / date / f"scientific-workflow-provenance__{rec}" / "collection.json"
    data = json.loads(path.read_text(encoding="utf-8"))
    commands = bug["verify_cmds"].splitlines()
    tests = bug["tests"]
    data.update({
        "sample_id": str(1000 + idx),
        "bug_id": f"scientific-workflow-provenance-{rec}",
        "task_type": bug["task_type"],
        "bug_category": bug["category"],
        "go_version": "golang:1.23;go.mod:1.23;GOTOOLCHAIN:auto",
        "repro_determinism": "deterministic",
        "user_query": (root / date / f"scientific-workflow-provenance__{rec}" / "prompt.txt").read_text(encoding="utf-8"),
        "verify_cmds": bug["verify_cmds"],
        "gold_root_cause": f"文件: {', '.join(bug['core_files'])} 符号: {', '.join(bug['core_symbols'])} 机制: {bug['mechanism_summary']}",
        "success_criteria": bug["success_criteria"],
        "coverage_contract": {
            "version": 1,
            "claims": [
                {"id": f"{rec}_claim_{n}", "description": f"{test} 覆盖一个独立行为", "test": test, "command": command, "paths": [bug["core_files"][n - 1]]}
                for n, (test, command) in enumerate(zip(tests, commands), 1)
            ],
        },
    })
    path.write_text(json.dumps(data, ensure_ascii=False, indent=2) + "\n", encoding="utf-8")
