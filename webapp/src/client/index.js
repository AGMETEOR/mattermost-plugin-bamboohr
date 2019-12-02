import axios from 'axios';

class Client {
    constructor() {
        this.axiosInstance = axios.create({
            baseUrl: '/plugins/bamboohr/api/v1',
        });
    }

    getEmployees = async () => {
        return this.doGet('/plugins/bamboohr/api/v1/employees');
    }

    doGet = async (url) => {
        try {
            const response = await this.axiosInstance.get(url);
            return response.data;
        } catch (error) {
            throw error;
        }
    }
}

const client = new Client();

export default client;