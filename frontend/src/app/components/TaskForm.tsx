import { useState } from 'react';

interface TaskFormProps {
  onSubmit: (title: string, description: string) => Promise<void>;
}

export const TaskForm = ({ onSubmit }: TaskFormProps) => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSubmit(title, description);
    setTitle('');
    setDescription('');
  };

  return (
    <div className="mb-8 p-4 bg-white rounded-lg shadow">
      <h2 className="text-xl font-semibold mb-4">新規タスク作成</h2>
      <form onSubmit={handleSubmit} className="space-y-4">
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
  );
};
