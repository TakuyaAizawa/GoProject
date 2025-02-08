import { useState, useEffect } from 'react';
import { Task } from '../types';
import { TaskService } from '../services/taskService';

export const useTasks = () => {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [editingTask, setEditingTask] = useState<Task | null>(null);

  const fetchTasks = async () => {
    try {
      const response = await TaskService.getTasks();
      setTasks(response);
    } catch (error) {
      console.error('タスクの取得中にエラーが発生しました:', error);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const createTask = async (title: string, description: string) => {
    try {
      await TaskService.createTask({ title, description });
      fetchTasks();
      return true;
    } catch (error) {
      console.error('タスクの作成中にエラーが発生しました:', error);
      return false;
    }
  };

  const updateTask = async (task: Task) => {
    try {
      await TaskService.updateTask(task);
      setEditingTask(null);
      fetchTasks();
      return true;
    } catch (error) {
      console.error('タスクの更新中にエラーが発生しました:', error);
      return false;
    }
  };

  const deleteTask = async (id: number) => {
    try {
      await TaskService.deleteTask(id);
      fetchTasks();
      return true;
    } catch (error) {
      console.error('タスクの削除中にエラーが発生しました:', error);
      return false;
    }
  };

  return {
    tasks,
    editingTask,
    setEditingTask,
    createTask,
    updateTask,
    deleteTask,
  };
};
