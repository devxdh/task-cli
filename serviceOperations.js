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

const getTask = async () => {
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


const AddTask = async (name) => {
    try {
        const data = dataFormat(name)
        const dataArray = await getTask()
        dataArray.push(data)
        const parsedArray = JSON.stringify(dataArray, null, 2)
        fs.writeFileSync(filepath, parsedArray)
    } catch (err) {
        console.error(`AddTask Error: ${err}`)
    }
}