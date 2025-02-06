'use client'

import { useState, useEffect } from 'react'
import axios from 'axios'
import Link from 'next/link'

interface Task {
  id: number
  title: string
  description: string
}

export default function TaskManager() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [editingTask, setEditingTask] = useState<Task | null>(null)
  const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';
  
  const fetchTasks = async () => {
    try {
      const response = await axios.get(`${API_URL}/api/tasks`);
      setTasks(response.data);
    } catch (error) {
      console.error('タスクの取得中にエラーが発生しました:', error);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const createTask = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      await axios.post(`${API_URL}/api/tasks`, {
        title,
        description
      })
      setTitle('')
      setDescription('')
      fetchTasks()
    } catch (error) {
      console.error('タスクの作成中にエラーが発生しました:', error)
    }
  }

  const startEdit = (task: Task) => {
    setEditingTask(task);
  };

  const updateTask = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!editingTask) return;

    try {
      await axios.put(`${API_URL}/api/task?id=${editingTask.id}`, {
        title: editingTask.title,
        description: editingTask.description
      });
      setEditingTask(null);
      fetchTasks();
    } catch (error) {
      console.error('タスクの更新中にエラーが発生しました:', error);
    }
  };

  const deleteTask = async (id: number) => {
    if (!confirm('このタスクを削除してもよろしいですか？')) return;
    
    try {
      await axios.delete(`${API_URL}/api/task?id=${id}`);
      fetchTasks();
    } catch (error) {
      console.error('タスクの削除中にエラーが発生しました:', error);
    }
  };

  return (
    <main className="max-w-4xl mx-auto p-4">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-2xl font-bold">タスク管理</h1>
        <Link 
          href="/tasks" 
          className="px-4 py-2 border rounded-md hover:bg-gray-100 transition-colors"
        >
          TODOリストへ→
        </Link>
      </div>
      
      {/* タスク作成フォーム */}
      <div className="mb-8 p-4 bg-white rounded-lg shadow">
        <h2 className="text-xl font-semibold mb-4">新規タスク作成</h2>
        <form onSubmit={createTask} className="space-y-4">
          <div>
            <label htmlFor="title" className="block text-sm font-medium text-gray-700">
              タスク名
            </label>
            <input
              id="title"
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="mt-1 w-full p-2 border rounded-md"
              placeholder="タスク名を入力してください"
              required
            />
          </div>
          <div>
            <label htmlFor="description" className="block text-sm font-medium text-gray-700">
              説明
            </label>
            <input
              id="description"
              type="text"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              className="mt-1 w-full p-2 border rounded-md"
              placeholder="説明を入力してください"
              required
            />
          </div>
          <button 
            type="submit"
            className="w-full bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 transition-colors"
          >
            タスクを追加
          </button>
        </form>
      </div>

      {/* タスク一覧 */}
      <div className="space-y-4">
        {tasks.map(task => (
          <div 
            key={task.id} 
            className="p-4 bg-white rounded-lg shadow"
          >
            {editingTask?.id === task.id ? (
              // 編集フォーム
              <form onSubmit={updateTask} className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700">タイトル</label>
                  <input
                    type="text"
                    value={editingTask.title}
                    onChange={(e) => setEditingTask({...editingTask, title: e.target.value})}
                    className="mt-1 w-full p-2 border rounded-md"
                    required
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700">説明</label>
                  <input
                    type="text"
                    value={editingTask.description}
                    onChange={(e) => setEditingTask({...editingTask, description: e.target.value})}
                    className="mt-1 w-full p-2 border rounded-md"
                    required
                  />
                </div>
                <div className="flex gap-2">
                  <button 
                    type="submit"
                    className="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600"
                  >
                    保存
                  </button>
                  <button 
                    type="button"
                    onClick={() => setEditingTask(null)}
                    className="px-4 py-2 border rounded-md hover:bg-gray-100"
                  >
                    キャンセル
                  </button>
                </div>
              </form>
            ) : (
              // タスク表示
              <div>
                <div className="flex justify-between items-start">
                  <div>
                    <h2 className="font-bold text-lg">{task.title}</h2>
                    <p className="text-gray-600 mt-2">{task.description}</p>
                  </div>
                  <div className="flex gap-2">
                    <button
                      onClick={() => startEdit(task)}
                      className="px-3 py-1 text-blue-600 hover:bg-blue-50 rounded-md"
                    >
                      編集
                    </button>
                    <button
                      onClick={() => deleteTask(task.id)}
                      className="px-3 py-1 text-red-600 hover:bg-red-50 rounded-md"
                    >
                      削除
                    </button>
                  </div>
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    </main>
  )
}