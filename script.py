import os
import sys
from pathlib import Path
from datetime import datetime

def explore_directory(directory_path, output_file):
    """
    ディレクトリ内のすべてのファイルを再帰的に探索し、
    ファイル名、パス、内容をファイルに出力します。
    
    Args:
        directory_path (str): 探索を開始するディレクトリのパス
        output_file (file): 出力先のファイルオブジェクト
    """
    try:
        # Pathオブジェクトに変換
        root_path = Path(directory_path)
        
        # ヘッダー情報を書き込み
        output_file.write(f'ファイル探索レポート\n')
        output_file.write(f'実行日時: {datetime.now().strftime("%Y-%m-%d %H:%M:%S")}\n')
        output_file.write(f'探索ディレクトリ: {root_path.absolute()}\n')
        output_file.write('='*50 + '\n\n')
        
        # ディレクトリ内のすべてのファイルを再帰的に取得
        for file_path in root_path.rglob('*'):
            if file_path.is_file():  # ディレクトリは除外
                try:
                    # ファイル情報の書き込み
                    output_file.write('\n' + '='*50 + '\n')
                    output_file.write(f'ファイル名: {file_path.name}\n')
                    output_file.write(f'絶対パス: {file_path.absolute()}\n')
                    output_file.write(f'相対パス: {file_path.relative_to(root_path)}\n')
                    output_file.write(f'サイズ: {file_path.stat().st_size:,} bytes\n')
                    output_file.write('-'*50 + '\n')
                    output_file.write('ファイル内容:\n')
                    
                    # テキストファイルの場合のみ内容を書き込み
                    try:
                        # バイナリファイルかどうかの簡易チェック
                        with open(file_path, 'rb') as f:
                            content = f.read(1024)  # 最初の1024バイトを読み込む
                            if b'\x00' in content:
                                output_file.write('[バイナリファイル - 内容は表示できません]\n')
                            else:
                                # テキストファイルとして読み込み
                                with open(file_path, 'r', encoding='utf-8') as f:
                                    output_file.write(f.read() + '\n')
                    except UnicodeDecodeError:
                        output_file.write('[バイナリファイルまたは未対応のエンコーディング - 内容は表示できません]\n')
                    except Exception as e:
                        output_file.write(f'[ファイル内容の読み取りエラー: {str(e)}]\n')
                
                except Exception as e:
                    output_file.write(f'[ファイル {file_path} の処理中にエラーが発生: {str(e)}]\n')
    
    except Exception as e:
        output_file.write(f'エラーが発生しました: {str(e)}\n')
        return

def main():
    # コマンドライン引数からディレクトリパスを取得
    if len(sys.argv) > 1:
        directory_path = sys.argv[1]
    else:
        # 引数がない場合は現在のディレクトリを使用
        directory_path = '.'
    
    # 出力ファイル名を生成（現在の日時を含む）
    output_filename = f'file_explorer_{datetime.now().strftime("%Y%m%d_%H%M%S")}.txt'
    
    print(f'ディレクトリ "{directory_path}" の探索を開始します...')
    print(f'結果は {output_filename} に出力されます。')
    
    # 出力ファイルをUTF-8で開く
    with open(output_filename, 'w', encoding='utf-8') as output_file:
        explore_directory(directory_path, output_file)
    
    print('探索が完了しました。')

if __name__ == '__main__':
    main()