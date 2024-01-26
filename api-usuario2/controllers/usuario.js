module.exports = () => {
    const controller = {};

    controller.usuario = (_, res) => res.status(200).json({ ola: 123456 });

    return controller;
}