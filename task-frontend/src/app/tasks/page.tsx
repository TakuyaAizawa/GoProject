'use client'

import { useState, useEffect } from 'react'
import Link from 'next/link'
import axios from 'axios'

interface Todo {
  id: number
  text: string
  created_at: string
  updated_at: string
  isEditing?: boolean
}

export default function Home() {
  const [todos, setTodos] = useState<Todo[]>([])
  const [newTodo, setNewTodo] = useState('')
  const [isLoading, setIsLoading] = useState(true)
  const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

  const fetchTodos = async () => {
    try {
        const url = `${API_URL}/api/todos`;
        console.log('Fetching todos from:', url);
        const response = await axios.get(url);
        // レスポンスのデータ構造を確認
        console.log('Response:', response);
        // response.dataが配列であることを確認
        const todos = Array.isArray(response.data) ? response.data : [];
        setTodos(todos.map((todo: Todo) => ({ ...todo, isEditing: false })));
    } catch (error) {
        console.error('TODOの取得中にエラーが発生しました:', error);
        setTodos([]); // エラー時は空配列をセット
    } finally {
        setIsLoading(false);
    }
};

  useEffect(() => {
    fetchTodos();
  }, []);

  const addTodo = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!newTodo.trim()) return;

    try {
      await axios.post(`${API_URL}/api/todos`, {
        text: newTodo.trim()
      });
      setNewTodo('');
      fetchTodos();
    } catch (error) {
      console.error('TODOの作成中にエラーが発生しました:', error);
    }
  }

  const startEdit = (id: number) => {
    setTodos(todos.map(todo => 
      todo.id === id ? { ...todo, isEditing: true } : todo
    ));
  }

  const updateTodo = async (id: number, newText: string) => {
    if (!newText.trim()) return;

    try {
      await axios.put(`${API_URL}/api/todo?id=${id}`, {
        text: newText.trim()
      });
      setTodos(todos.map(todo =>
        todo.id === id ? { ...todo, isEditing: false } : todo
      ));
      fetchTodos();
    } catch (error) {
      console.error('TODOの更新中にエラーが発生しました:', error);
    }
  }

  const cancelEdit = (id: number) => {
    setTodos(todos.map(todo =>
      todo.id === id ? { ...todo, isEditing: false } : todo
    ));
  }

  const removeTodo = async (id: number) => {
    if (!confirm('このTODOを削除してもよろしいですか？')) return;

    try {
      await axios.delete(`${API_URL}/api/todo?id=${id}`);
      fetchTodos();
    } catch (error) {
      console.error('TODOの削除中にエラーが発生しました:', error);
    }
  }

  if (isLoading) {
    return (
      <main className="max-w-2xl mx-auto p-4">
        <div className="flex justify-center items-center h-32">
          読み込み中...
        </div>
      </main>
    );
  }

  return (
    <main className="max-w-2xl mx-auto p-4">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-2xl font-bold">TODOリスト</h1>
        <Link 
          href="/" 
          className="px-4 py-2 border rounded-md hover:bg-gray-100 transition-colors"
        >
          ← タスク管理へ戻る
        </Link>
      </div>
      
      <form onSubmit={addTodo} className="mb-6 flex gap-2">
        <input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
          placeholder="新しいTODOを入力..."
          className="flex-1 p-2 border rounded-md"
        />
        <button 
          type="submit"
          className="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 transition-colors"
        >
          追加
        </button>
      </form>

      <ul className="space-y-2">
        {todos.map((todo) => (
          <li 
            key={todo.id} 
            className="p-3 bg-white rounded-lg shadow"
          >
            {todo.isEditing ? (
              <div className="flex gap-2">
                <input
                  type="text"
                  defaultValue={todo.text}
                  className="flex-1 p-2 border rounded-md"
                  onKeyDown={(e) => {
                    if (e.key === 'Enter') {
                      updateTodo(todo.id, e.currentTarget.value)
                    } else if (e.key === 'Escape') {
                      cancelEdit(todo.id)
                    }
                  }}
                  autoFocus
                />
                <button
                  onClick={() => cancelEdit(todo.id)}
                  className="px-3 py-1 border rounded-md hover:bg-gray-100"
                >
                  キャンセル
                </button>
              </div>
            ) : (
              <div className="flex items-center justify-between">
                <span className="flex-1">{todo.text}</span>
                <div className="flex gap-2">
                  <button
                    onClick={() => startEdit(todo.id)}
                    className="px-3 py-1 text-blue-600 hover:bg-blue-50 rounded-md"
                  >
                    編集
                  </button>
                  <button
                    onClick={() => removeTodo(todo.id)}
                    className="px-3 py-1 text-red-600 hover:bg-red-50 rounded-md"
                  >
                    削除
                  </button>
                </div>
              </div>
            )}
          </li>
        ))}
      </ul>

      {todos.length > 0 && (
        <p className="mt-4 text-sm text-gray-600">
          ヒント: 編集中はEnterキーで保存、Escキーでキャンセルできます
        </p>
      )}
    </main>
  )
}