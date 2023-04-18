import { useSelector } from 'react-redux';
import type { RootState } from '@/store/store';

export default function num() {
  const count = useSelector((state: RootState) => state.counter.value);
  return (
    <div>
      <p>Count: {count}</p>
    </div>
  );
}
