module.exports = () => {
    const controller = {};

    controller.usuario = (req, res) => {
        console.log('[API 2] - Acessando dados de usuário');

        for (const [key, value] of Object.entries(req.headers)) {
            console.log(`[API 2] Header: ${key} = ${value}`);
        }

        return res.status(200).json({ nome: "TigreDoMexico - API 2" })
    } ;

    return controller;
}