import axios from 'axios'
import authorization from './authorization'
import friends from './friends'

export const axiosInstance = axios.create({
  baseURL: 'http://localhost:12001',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json'
  }
})

export const API = {
  authorization,
  friends,
}
