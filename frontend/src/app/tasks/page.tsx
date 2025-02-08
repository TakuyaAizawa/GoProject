'use client'

import { useState } from 'react'
import Link from 'next/link'
import { TodoList } from '../components/TodoList'
import { useTodos } from '../hooks/useTodos'

export default function TodoPage() {
  const {
    todos,
    isLoading,
    setTodos,
    addTodo,
    updateTodo,
    deleteTodo,
  } = useTodos();
  const [newTodo, setNewTodo] = useState('');

  const handleAddTodo = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!newTodo.trim()) return;
    await addTodo(newTodo);
    setNewTodo('');
  };

  const startEdit = (id: number) => {
    setTodos(todos.map(todo => 
      todo.id === id ? { ...todo, isEditing: true } : todo
    ));
  };

  const cancelEdit = (id: number) => {
    setTodos(todos.map(todo =>
      todo.id === id ? { ...todo, isEditing: false } : todo
    ));
  };

  const handleDeleteTodo = async (id: number) => {
    if (!confirm('このTODOを削除してもよろしいですか？')) return;
    await deleteTodo(id);
  };

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
      
      <form onSubmit={handleAddTodo} className="mb-6 flex gap-2">
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

      <TodoList
        todos={todos}
        onStartEdit={startEdit}
        onUpdate={updateTodo}
        onCancelEdit={cancelEdit}
        onDelete={handleDeleteTodo}
      />

      {todos.length > 0 && (
        <p className="mt-4 text-sm text-gray-600">
          ヒント: 編集中はEnterキーで保存、Escキーでキャンセルできます
        </p>
      )}
    </main>
  );
}
