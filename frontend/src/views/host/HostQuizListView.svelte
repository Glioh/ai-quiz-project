<script lang="ts">
    import QuizCard from '../../lib/QuizCard.svelte';
    import Button from '../../lib/Button.svelte';
    import GenerateQuiz from '../../lib/QuizGeneratorForm.svelte';
    import QuizCreationChoice from '../../lib/QuizCreationChoice.svelte';
    import type { Quiz } from '../../model/quiz';
    import { apiService } from '../../service/api';

    let quizzes: Quiz[] = [];
    let showForm = false;
    let showCreationChoice = false;
    let creationMode: 'ai' | 'manual' | null = null;

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

    function handleQuizDeleted(event: CustomEvent<string>) {
        const deletedQuizId = event.detail;
        quizzes = quizzes.filter(quiz => quiz.id !== deletedQuizId);
    }

    function toggleCreationChoice() {
        showCreationChoice = !showCreationChoice;
        creationMode = null;
        showForm = false;
    }

    async function handleCreationChoice(event: CustomEvent<'ai' | 'manual'>) {
        creationMode = event.detail;
        showCreationChoice = false;
        
        if (creationMode === 'ai') {
            showForm = true;
        } else {
            // Create a blank quiz and redirect to edit page
            try {
                const newQuiz = await apiService.createQuiz('New Quiz');
                window.location.href = `/#/edit/${newQuiz.id}`;
            } catch (error) {
                console.error('Failed to create quiz:', error);
                alert('Failed to create quiz');
            }
        }
    }

    async function handleQuizGenerated() {
        showForm = false;
        creationMode = null;
        await getQuizzes();
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
        <h2 class="text-4xl font-bold">
            {#if showForm}
                Create AI Quiz
            {:else if showCreationChoice}
                Create New Quiz
            {:else}
                Your Quizzes
            {/if}
        </h2>
        <div class="flex gap-2">
            {#if !showForm && !showCreationChoice}
                <Button on:click={toggleCreationChoice}>New Quiz</Button>
            {:else}
                <Button on:click={() => {
                    showForm = false;
                    showCreationChoice = false;
                    creationMode = null;
                }}>Cancel</Button>
            {/if}
            <Button on:click={goBack}>Back to Home</Button>
        </div>
    </div>
    
    {#if showCreationChoice}
        <QuizCreationChoice on:choice={handleCreationChoice} />
    {:else if showForm}
        <GenerateQuiz on:generated={handleQuizGenerated} />
    {:else}
        <div class="flex flex-col gap-2 mt-4">
            {#each quizzes as quiz (quiz.id)}
                <QuizCard
                    {quiz}
                    on:host
                    on:deleted={handleQuizDeleted}
                />
            {/each}
        </div>
    {/if}
</div>