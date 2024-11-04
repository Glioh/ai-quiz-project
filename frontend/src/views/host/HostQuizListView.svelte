<script lang="ts">
    import QuizCard from '../../lib/QuizCard.svelte';
    import Button from '../../lib/Button.svelte';
import type{Quiz} from '../../model/quiz';
import NewQuizButton from '../../lib/edit/NewQuizButton.svelte';
    import { apiService } from '../../service/api';

    let quizzes: Quiz[] = [];

    async function getQuizzes(): Promise<Quiz[]> {
        let response = await fetch('http://localhost:3000/api/quizzes');
        if (!response.ok) {
            alert("Failed to fetch quizzes!");
            return [];
        }

        let json = await response.json();
        return json;
    }

    function goBack() {
        window.location.href = "/";
    }

    (async function() {
        quizzes = await apiService.getQuizzes();
    })();
</script>

<div class="p-8">
    <div class="flex justify-between items-center">
        <h2 class="text-4xl font-bold">Your quizzes</h2>
        <div class="flex gap-2">
            <NewQuizButton />
            <Button on:click={goBack}>Back to Home</Button>
        </div>
        <Button on:click={goBack}>Back to Home</Button>
    </div>
    <div class="flex flex-col gap-2 mt-4">
        {#each quizzes as quiz (quiz.id)}
            <QuizCard on:host {quiz} />
        {/each}
    </div>
</div>