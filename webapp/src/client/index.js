import axios from 'axios';

class Client {
    constructor(){
        this.axiosInstance = axios.create({
            baseUrl: "/plugins/bamboohr"
        })
    }

    getEmployees = async() => {
        return this.doGet("/employees")
    }

    doGet = async(url) => {
        try {
            const response = await this.axiosInstance.get(url)
            return response.body  
        } catch (error) {
            throw error  
        }
    }
}

const client = new Client();

export default client;