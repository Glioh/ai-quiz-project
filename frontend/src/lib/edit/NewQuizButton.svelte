<script lang="ts">
    import Button from "../Button.svelte";
    import { apiService } from "../../service/api";

    async function createNewQuiz() {
        const quiz = await apiService.createQuiz("New Quiz");
        window.location.href = `/#/edit/${quiz.id}`;
    }
</script>

<Button on:click={createNewQuiz}>
    Create New Quiz
</Button>