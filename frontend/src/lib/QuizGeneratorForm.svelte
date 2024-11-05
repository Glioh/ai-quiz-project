<script lang="ts">
    import { apiService } from "../service/api";
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();
    let name = '';
    let prompt = '';
    let isLoading = false;
    let errorMessage = '';

    async function generateQuiz(event: Event) {
        event.preventDefault();
        errorMessage = '';
        
        if (!name.trim() || !prompt.trim()) {
            errorMessage = 'Both name and prompt are required';
            return;
        }

        try {
            isLoading = true;
            console.log('Generating quiz with:', { name, prompt });
            
            const quiz = await apiService.generateAIQuiz(name.trim(), prompt.trim());
            console.log('Generated quiz:', quiz);
            
            if (!quiz.id) {
                throw new Error('Invalid quiz response - missing ID');
            }

            dispatch('generated', { quiz });
            window.location.href = `/quiz/${quiz.id}`;
        } catch (error) {
            console.error('Failed to generate quiz:', error);
            errorMessage = error instanceof Error ? error.message : 'Failed to generate quiz. Please try again.';
        } finally {
            isLoading = false;
        }
    }
</script>

<div class="max-w-md mx-auto mt-8 p-6 bg-white rounded-lg shadow">
    <h2 class="text-2xl font-bold mb-4">Generate AI Quiz</h2>
    
    {#if errorMessage}
        <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {errorMessage}
        </div>
    {/if}

    <form on:submit={generateQuiz}>
        <div class="mb-4">
            <label for="quiz-name" class="block text-gray-700 mb-2">Quiz Name</label>
            <input
                id="quiz-name"
                type="text"
                bind:value={name}
                class="w-full p-2 border rounded"
                required
            />
        </div>
        <div class="mb-4">
            <label for="quiz-prompt" class="block text-gray-700 mb-2">Topic/Prompt</label>
            <textarea
                id="quiz-prompt"
                bind:value={prompt}
                class="w-full p-2 border rounded h-32"
                placeholder="Describe the topic or subject for your quiz..."
                required
            />
        </div>
        <button
            type="submit"
            disabled={isLoading}
            class="w-full bg-blue-500 text-white p-2 rounded hover:bg-blue-600 disabled:bg-gray-400"
        >
            {isLoading ? 'Generating...' : 'Generate Quiz'}
        </button>
    </form>
</div>