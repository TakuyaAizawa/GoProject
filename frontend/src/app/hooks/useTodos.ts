import { useState, useEffect } from 'react';
import { Todo } from '../types';
import { TodoService } from '../services/todoService';

export const useTodos = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  const fetchTodos = async () => {
    try {
      const response = await TodoService.getTodos();
      setTodos(response.map((todo: Todo) => ({ ...todo, isEditing: false })));
    } catch (error) {
      console.error('TODOの取得中にエラーが発生しました:', error);
      setTodos([]);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  const addTodo = async (text: string): Promise<void> => {
    try {
      await TodoService.createTodo(text);
      fetchTodos();
    } catch (error) {
      console.error('TODOの作成中にエラーが発生しました:', error);
    }
  };

  const updateTodo = async (id: number, text: string): Promise<void> => {
    try {
      await TodoService.updateTodo(id, text);
      setTodos(todos.map(todo =>
        todo.id === id ? { ...todo, isEditing: false } : todo
      ));
      fetchTodos();
    } catch (error) {
      console.error('TODOの更新中にエラーが発生しました:', error);
    }
  };

  const deleteTodo = async (id: number): Promise<void> => {
    try {
      await TodoService.deleteTodo(id);
      fetchTodos();
    } catch (error) {
      console.error('TODOの削除中にエラーが発生しました:', error);
    }
  };

  return {
    todos,
    isLoading,
    setTodos,
    addTodo,
    updateTodo,
    deleteTodo,
  };
};
