# Page 7

## Text Content

```
<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
<script src="https://cdn.tailwindcss.com"></script>
<div class="container ml-40 my-40">
<script>
$(document).ready(function () {
// Toggle dropdown when clicking on the button
$("#hs-pro-inthmtdid1").on("click", function (e) {
e.stopPropagation();
// Toggle classes that control visibility and opacity
$(this)
.siblings(".hs-dropdown-menu")
.toggleClass("hidden opacity-0");
});
// Hide the dropdown if a click occurs outside the dropdown container
$(document).on("click", function (e) {
if (!$(e.target).closest(".hs-dropdown").length) {
$(".hs-dropdown-menu").addClass("hidden opacity-0");
}
});
});
</script>
<div class="hs-dropdown [--auto-close:inside] relative inline-flex">
<button
id="hs-pro-inthmtdid1"
type="button"
class="py-1.5 lg:px-2.5 flex justify-center items-center text-[13px] text-g
>
<span
class="shrink-0 size-2 block me-2 bg-blue-100 rounded-full"
></span>
<span class="max-w-[5rem] truncate">
<svg
fill="#000000"
width="24"
height="24"
viewBox="0 0 256 256"
id="Flat"
xmlns="http://www.w3.org/2000/svg"
>
<path


```

