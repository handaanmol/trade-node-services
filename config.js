var util = require('util');
var path = require('path');
var hfc = require('fabric-client');

var file = 'network-config%s.yaml';

var env = process.env.TARGET_NETWORK;
if (env)
	file = util.format(file, '-' + env);
else
	file = util.format(file, '');
// indicate to the application where the setup file is located so it able
// to have the hfc load it to initalize the fabric client instance
hfc.setConfigSetting('network-connection-profile-path',path.join(__dirname, 'artifacts' ,file));
hfc.setConfigSetting('IM1-connection-profile-path',path.join(__dirname, 'artifacts', 'org1.yaml'));
hfc.setConfigSetting('EB1-connection-profile-path',path.join(__dirname, 'artifacts', 'org2.yaml'));
hfc.setConfigSetting('EB2-connection-profile-path',path.join(__dirname, 'artifacts', 'org3.yaml'));
hfc.setConfigSetting('AB1-connection-profile-path',path.join(__dirname, 'artifacts', 'org4.yaml'));
hfc.setConfigSetting('AB2-connection-profile-path',path.join(__dirname, 'artifacts', 'org5.yaml'));
hfc.setConfigSetting('CB1-connection-profile-path',path.join(__dirname, 'artifacts', 'org6.yaml'));
// some other settings the application might need to know
hfc.addConfigFile(path.join(__dirname, 'config.json'));
