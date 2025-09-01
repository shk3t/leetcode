return {
  debug = {
    go = {
      {
        name = "Leetcode debug",
        type = "delve",
        request = "launch",
        program = "main.go"
      },
    },
  },
  run = {
    go = "go run main.go",
  },
}
