const app = require('./express')();
const port = app.get('port');

app.listen(port, () => {
  console.log(`[API 2] - Servidor rodando na porta ${port}`)
});