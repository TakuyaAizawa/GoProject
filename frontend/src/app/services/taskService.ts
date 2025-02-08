import axios from 'axios';
import { Task } from '../types';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export const TaskService = {
  getTasks: async () => {
    const response = await axios.get(`${API_URL}/api/tasks`);
    return response.data;
  },

  createTask: async (task: Omit<Task, 'id'>) => {
    return axios.post(`${API_URL}/api/tasks`, task);
  },

  updateTask: async (task: Task) => {
    return axios.put(`${API_URL}/api/task?id=${task.id}`, {
      title: task.title,
      description: task.description
    });
  },

  deleteTask: async (id: number) => {
    return axios.delete(`${API_URL}/api/task?id=${id}`);
  }
};
