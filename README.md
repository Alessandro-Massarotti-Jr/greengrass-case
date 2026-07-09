# greengrass-case

<p>
  <img src="https://img.shields.io/badge/made%20by-Alessandro%20Massarotti%20Jr-488e29?style=flat-square">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/Alessandro-Massarotti-Jr/greengrass-case?color=488e29&style=flat-square">
  <img alt="GitHub Top Language" src="https://img.shields.io/github/languages/top/Alessandro-Massarotti-Jr/greengrass-case?color=488e29&style=flat-square">
</p>

Estudo de casos de uso do AWS GreenGrass


Aqui eu quero o seguinte:
- uma maquina para simular um dispositivo Edge
- Duas aplicações distintas para serem um componente do greengrass cada uma delas
- Uma action automatica para adicionar novas versões dos componentes ao greengrass com pushs na main

Oque as aplicações fazem não é tão relevante agora, o Objetivo principal é testar o fluxo do greengrass como monitoramento e deploy de versões


Passo a passo do GreenGrass:

1. Configurar um Core Device
   1. instalar Java no dispositivo
   2. Criar um usuario no IAM para ter acesso aos recursos do GreenGrass
   3. Adicionar as credenciais deste usuario no dispositivo desejado
   4. instalar o greengrass no dispositivo
      1. curl -s https://d2s8p88vqu9w66.cloudfront.net/releases/greengrass-nucleus-latest.zip > greengrass-nucleus-latest.zip && unzip greengrass-nucleus-latest.zip -d GreengrassInstaller
      2. sudo -E java -Droot="/greengrass/v2" -Dlog.store=FILE -jar ./GreengrassInstaller/lib/Greengrass.jar --aws-region us-east-1 --thing-name unidade-0  --component-default-user ggc_user:ggc_group --provision true --setup-system-service true --deploy-dev-tools true
   5. Adicionar a integração em um bucket S3
   6. Criar um componente greengrass para aquele bucket
   7. Criar um deployment para o grupo de things

---

Developed by [Alessandro Massarotti Jr](https://github.com/Alessandro-Massarotti-Jr) 🤖
