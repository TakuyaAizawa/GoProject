#!/bin/bash

# useTodos.tsを更新
cat > src/app/hooks/useTodos.ts << 'EOL'
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
EOL

# TodoList.tsxのPropsの型定義を確認用に表示
echo "TodoList.tsx の型定義:"
cat > src/app/components/TodoList.tsx << 'EOL'
import { Todo } from '../types';

interface TodoListProps {
  todos: Todo[];
  onStartEdit: (id: number) => void;
  onUpdate: (id: number, text: string) => Promise<void>;
  onCancelEdit: (id: number) => void;
  onDelete: (id: number) => Promise<void>;
}

export const TodoList = ({
  todos,
  onStartEdit,
  onUpdate,
  onCancelEdit,
  onDelete,
}: TodoListProps) => {
  return (
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
                    onUpdate(todo.id, e.currentTarget.value)
                  } else if (e.key === 'Escape') {
                    onCancelEdit(todo.id)
                  }
                }}
                autoFocus
              />
              <button
                onClick={() => onCancelEdit(todo.id)}
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
                  onClick={() => onStartEdit(todo.id)}
                  className="px-3 py-1 text-blue-600 hover:bg-blue-50 rounded-md"
                >
                  編集
                </button>
                <button
                  onClick={() => onDelete(todo.id)}
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
  );
};
EOL

echo "型の定義を更新しました。"
echo "更新されたファイル："
echo "- src/app/hooks/useTodos.ts"
echo "- src/app/components/TodoList.tsx"