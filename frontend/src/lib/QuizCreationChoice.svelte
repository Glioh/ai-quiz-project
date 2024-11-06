<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import Button from './Button.svelte';

    const dispatch = createEventDispatcher();

    function handleAIGeneration() {
        dispatch('choice', 'ai');
    }

    function handleManualCreation() {
        dispatch('choice', 'manual');
    }
</script>

<div class="bg-white p-6 rounded-lg shadow-lg max-w-md mx-auto mt-8">
    <h3 class="text-2xl font-bold mb-4 text-center">Create New Quiz</h3>
    
    <div class="space-y-6">
        <!-- svelte-ignore a11y-no-redundant-roles -->
        <button type="button" class="p-4 border rounded-lg hover:bg-gray-50 cursor-pointer" on:click={handleAIGeneration} role="button">
            <h4 class="text-lg font-semibold mb-2">AI Quiz Generation</h4>
            <p class="text-gray-600">Let AI generate quiz questions based on your prompt</p>
        </button>

        <!-- svelte-ignore a11y-no-redundant-roles -->
        <button type="button" class="p-4 border rounded-lg hover:bg-gray-50 cursor-pointer" on:click={handleManualCreation} role="button">
            <h4 class="text-lg font-semibold mb-2">Manual Creation</h4>
            <p class="text-gray-600">Create your own quiz questions from scratch</p>
        </button>
    </div>
</div>