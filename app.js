const { DisplayTask, AddTask, UpdateTask, DeleteTask } = require("./serviceOperations")

const userInput = process.argv

const command = userInput[2]
const taskName = userInput[3]

const run = async () => {
    switch (command) {
        case "get":
            await DisplayTask()
            break;
        case "add":
            await AddTask(taskName)
            break;
        case "update":
            await UpdateTask(taskName)
            break;
        case "delete":
            await DeleteTask(taskName)
            break;
    }
}

run() 