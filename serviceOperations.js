const fs = require("node:fs/promises")
const filepath = "./db.json"
const crypto = require('node:crypto')

const dataFormat = {
    id: crypto.randomUUID(),
    name: taskName,
    status: false
}

const existingDataHandler = async () => {
    try {
        const JSONFileData = await fs.readFile(filepath, 'utf8')
        const JSONFileArray = await JSON.parse(JSONFileData)
        return JSONFileArray;
    } catch (err) {
        console.error(`Error occured in existingJSONFileHandler: ${err}`)
    }
}
const AddTask = async (name) => {
    try {

    } catch (err) {

    }
}