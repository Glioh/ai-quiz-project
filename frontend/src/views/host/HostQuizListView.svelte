<script lang="ts">
    import QuizCard from '../../lib/QuizCard.svelte';
    import Button from '../../lib/Button.svelte';
    import NewQuizButton from '../../lib/edit/NewQuizButton.svelte';
    import GenerateQuiz from '../../lib/QuizGeneratorForm.svelte';
    import type { Quiz } from '../../model/quiz';
    import { ApiService } from '../../service/api';

    let quizzes: Quiz[] = [];
    let showForm = false; // controls display of the form

    async function getQuizzes(): Promise<void> {
        try {
            const response = await fetch('http://localhost:3000/api/quizzes');
            if (!response.ok) {
                alert("Failed to fetch quizzes!");
                return;
            }
            quizzes = await response.json();
        } catch (error) {
            console.error("Failed to fetch quizzes:", error);
        }
    }

    function toggleForm() {
        showForm = !showForm;
    }

    async function handleQuizGenerated() {
        showForm = false; // hide form after quiz generation
        await getQuizzes(); // refresh the quiz list
    }

    function goBack() {
        window.location.href = "/";
    }

    (async function() {
        await getQuizzes();
    })();
</script>

<div class="p-8">
    <div class="flex justify-between items-center">
        <h2 class="text-4xl font-bold">{showForm ? 'Create a New Quiz' : 'Your Quizzes'}</h2>
        <div class="flex gap-2">
            <Button on:click={toggleForm}>{showForm ? 'Cancel' : 'New Quiz'}</Button>
            <Button on:click={goBack}>Back to Home</Button>
        </div>
    </div>
    
    {#if showForm}
        <!-- Quiz Generation Form -->
        <GenerateQuiz on:generated={handleQuizGenerated} />
    {:else}
        <!-- List of Quizzes -->
        <div class="flex flex-col gap-2 mt-4">
            {#each quizzes as quiz (quiz.id)}
                <QuizCard on:host {quiz} />
            {/each}
        </div>
    {/if}
</div>