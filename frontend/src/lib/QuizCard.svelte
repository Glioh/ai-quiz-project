<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import type { Quiz } from '../model/quiz';
    import Button from './Button.svelte';
    import { apiService } from '../service/api';

    export let quiz: Quiz;
    const dispatch = createEventDispatcher();

    async function handleDelete() {
        if (confirm('Are you sure you want to delete this quiz?')) {
            try {
                await apiService.deleteQuiz(quiz.id);
                dispatch('deleted', quiz.id);
            } catch (error) {
                console.error('Failed to delete quiz:', error);
                alert('Failed to delete quiz');
            }
        }
    }

    function handleHost() {
        dispatch('host', quiz);
    }

    function handleEdit() {
    window.location.href = `/#/edit/${quiz.id}`;
    }
</script>

<div class="bg-white shadow rounded-lg p-6">
    <div class="flex justify-between items-center">
        <h3 class="text-xl font-bold">{quiz.name}</h3>
        <div class="flex gap-2">
            <Button on:click={handleHost}>Host</Button>
            <Button on:click={handleEdit}>Edit</Button>
            <Button on:click={handleDelete} class="bg-red-500 hover:bg-red-600">Delete</Button>
        </div>
    </div>
    <p class="mt-2 text-gray-600">{quiz.questions?.length || 0} questions</p>
</div>