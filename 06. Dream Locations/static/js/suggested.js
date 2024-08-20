const suggestedSection = suggestedList.parentNode;

suggestedList.addEventListener("htmx:afterRequest", function () {
  const length = suggestedList.children.length;
  if (length === 0) {
    suggestedSection.remove();
  }
});
