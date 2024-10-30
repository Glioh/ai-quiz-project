<script lang="ts">
    import type { QuizChoice } from '../../model/quiz'; // Adjust the import path as necessary
    
    export let choice: QuizChoice;
    export let color: string;
</script>

<div CLASS="w-1/2 text-white font-bold flex items-center text-3x1 pl-14 {color}">
    {choice.name}
</div>