<script lang="ts">
    import type { QuizChoice } from '../../model/quiz'; // Adjust the import path as necessary
    
    export let color: string;
</script>

<div CLASS="w-1/2 text-white font-bold flex items-center text-5xl pl-14 {color}">
    <slot />
</div>