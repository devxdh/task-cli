const fs = require("node:fs")
const crypto = require('node:crypto')
const filepath = "./db.json"

const dataFormat = (name) => {
    const obj = {
        id: crypto.randomUUID(),
        name: name,
        status: false
    }
    return obj;
}

const GetTask = async () => {
    try {
        const content = fs.readFileSync(filepath, 'utf-8');
        if (!content.trim()) return [];

        const parsed = JSON.parse(content)
        return Array.isArray(parsed) ? parsed : [parsed];
    } catch (err) {
        console.error(`getTask: ${err}`);
        return [];
    }
}

const DisplayTask = async () => {
    try {
        await GetTask().then((e) => {
            if (e.length === 0) {
                console.log(`No Tasks`);
            } else {
                console.log(e);
            }
        })
    } catch (err) {
        console.error(`DisplayTask Error: ${err}`);
    }
}

const AddTask = async (taskName) => {
    try {
        const data = dataFormat(taskName)
        const dataArray = await GetTask()
        dataArray.push(data)
        const parsedArray = JSON.stringify(dataArray, null, 2)
        fs.writeFileSync(filepath, parsedArray)
        console.log(`Added Task Successfully ${taskName}`);
    } catch (err) {
        console.error(`AddTask Error: ${err}`)
    }
}

const UpdateTask = async (taskName) => {
    try {
        const dataArray = await GetTask();
        if (dataArray.length === 0) {
            console.log("There is no data in db.json file");
            return
        }
        const taskToUpdate = dataArray.find(task => task.name === taskName)

        if (taskToUpdate) {
            if (taskToUpdate.status == true) {
                console.log(`Task already updated: ${taskName}`);
                return
            }
            taskToUpdate.status = true;
            fs.writeFileSync(filepath, JSON.stringify(dataArray, null, 2))
            console.log(`Task Successfully Updated`)
        } else {
            console.log(`No Task Found: ${taskName}`);
            return
        }
    } catch (err) {
        console.error(`UpdateTask Error: ${err}`);
    }
}

const DeleteTask = async (taskName) => {
    try {
        const dataArray = await GetTask();
        if (dataArray.length === 0) {
            console.log("There is no data in db.json file");
            return
        }

        const updatedArray = dataArray.filter(task => task.name !== taskName)
        if (updatedArray.length === dataArray.length) {
            console.log(`No Task found to delete with name ${taskName}`);
            return
        }
        fs.writeFileSync(filepath, JSON.stringify(updatedArray, null, 2))
        console.log(`Task ${taskName} Deleted Successfully`);
    } catch (err) {
        console.error(`DeleteTask Error: ${err}`);

    }
}

module.exports = { DisplayTask, AddTask, UpdateTask, DeleteTask }