using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace ApiUsuario.Controllers;

[ApiController]
[Route("usuario")]
public class UsuarioController : ControllerBase
{
    private new const string Response = "TigreDoMexico - API 1";
    private readonly ILogger<UsuarioController> _logger;

    public UsuarioController(ILogger<UsuarioController> logger) => _logger = logger;

    [HttpGet]
    public string Index()
    {
        _logger.LogInformation("[API 1] - Acessando dados de usuário");
        this.ImprimirHeaders();
        
        return Response;
    }

    private void ImprimirHeaders()
    {
        foreach (var (key, value) in Request.Headers)
        {
            _logger.LogInformation("[API 1] Header: {Key} = {Value}", key, value);
        }
    }
}