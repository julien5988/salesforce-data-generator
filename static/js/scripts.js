document.getElementById('generate-btn').addEventListener('click', function() {
  fetch('/generate')
      .then(response => response.text())
      .then(data => {
          document.getElementById('response').textContent = data;
      })
      .catch(error => {
          console.error('Error:', error);
      });
});
