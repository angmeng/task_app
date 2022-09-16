import { writable, derived } from 'svelte/store';

export const apiData = writable({});

export const taskList = derived(apiData, ($taskResponse) => {
  if ($taskResponse.data){
    return $taskResponse.data.map(task => task);
  }
  return [];
});

export const meta = derived(apiData, ($taskResponse) => {
  if ($taskResponse.meta){
    return $taskResponse.meta;
  }
  return {};
});