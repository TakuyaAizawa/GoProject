import axios from 'axios';
import { Todo } from '../types';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export const TodoService = {
  getTodos: async () => {
    const response = await axios.get(`${API_URL}/api/todos`);
    return Array.isArray(response.data) ? response.data : [];
  },

  createTodo: async (text: string) => {
    return axios.post(`${API_URL}/api/todos`, { text: text.trim() });
  },

  updateTodo: async (id: number, text: string) => {
    return axios.put(`${API_URL}/api/todo?id=${id}`, { text: text.trim() });
  },

  deleteTodo: async (id: number) => {
    return axios.delete(`${API_URL}/api/todo?id=${id}`);
  }
};
