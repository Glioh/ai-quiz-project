<script lang="ts">
    import { apiService } from "../service/api";
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();
    let name = '';
    let prompt = '';
    let numQuestions = 10; // Default value
    let isLoading = false;
    let errorMessage = '';

    async function generateQuiz(event: Event) {
        event.preventDefault();
        errorMessage = '';
        
        // Check if all fields are filled
        if (!name.trim()) {
            errorMessage = 'Quiz name is required';
            return;
        }

        if (!prompt.trim()) {
            errorMessage = 'Topic/Prompt is required';
            return;
        }

        // Validate number of questions
        if (!numQuestions || numQuestions < 1 || numQuestions > 15) {
            errorMessage = 'Number of questions must be between 1 and 15';
            return;
        }

        try {
            isLoading = true;
            console.log('Generating quiz with:', { name, prompt, numQuestions });
            
            const quiz = await apiService.generateAIQuiz(name.trim(), prompt.trim(), numQuestions);
            console.log('Generated quiz:', quiz);
            
            if (!quiz.id) {
                throw new Error('Invalid quiz response - missing ID');
            }

            dispatch('generated', { quiz });
            window.location.href = `/#/edit/${quiz.id}`;
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
    
    <!-- Changed from blue to purple for notification banner -->
    <div class="mb-4 p-3 bg-purple-50 text-purple-700 rounded border border-purple-200">
        <p>⚠️ Please note: Quiz generation may take up to 1 minute to complete.</p>
    </div>

    {#if errorMessage}
        <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
            {errorMessage}
        </div>
    {/if}

    <form on:submit={generateQuiz}>
        <div class="mb-4">
            <label for="quiz-name" class="block text-gray-700 mb-2">Quiz Name *</label>
            <input
                id="quiz-name"
                type="text"
                bind:value={name}
                class="w-full p-2 border rounded"
                required
                placeholder="Enter quiz name"
            />
        </div>
        <div class="mb-4">
            <label for="num-questions" class="block text-gray-700 mb-2">Number of Questions (1-15) *</label>
            <input
                id="num-questions"
                type="number"
                bind:value={numQuestions}
                min="1"
                max="15"
                class="w-full p-2 border rounded"
                required
                on:input={(e) => {
                    const value = parseInt(e.currentTarget.value);
                    if (value < 1) numQuestions = 1;
                    if (value > 15) numQuestions = 15;
                }}
            />
        </div>
        <div class="mb-4">
            <label for="quiz-prompt" class="block text-gray-700 mb-2">Topic/Prompt *</label>
            <textarea
                id="quiz-prompt"
                bind:value={prompt}
                class="w-full p-2 border rounded h-32"
                placeholder="Describe the topic or subject for your quiz..."
                required
            />
        </div>
        <!-- Updated button with purple theme and enhanced styling -->
        <button
            type="submit"
            disabled={isLoading}
            class="w-full bg-purple-500 text-white p-2 rounded hover:bg-purple-600 disabled:bg-gray-400 transition-colors duration-200 shadow-sm hover:shadow-md focus:outline-none focus:ring-2 focus:ring-purple-400 focus:ring-opacity-50"
        >
            {isLoading ? 'Generating...' : 'Generate Quiz'}
        </button>
    </form>
</div>