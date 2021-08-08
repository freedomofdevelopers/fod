import requests

output_format = '''
[SwitchyOmega Conditions]
@with result

%s

* +direct
'''

url = 'https://raw.githubusercontent.com/freedomofdevelopers/fod/master/domains'


result = requests.get(url)
domains = result.text.split('\n')

output_file = open('out.txt', 'w')
output_domain_rules = []

for domain in domains:
    output_domain_rules.append((domain + ' +proxy\n'))

output_domain_rules[-2] = output_domain_rules[-2][:-1]

out = output_format % (''.join(output_domain_rules[:-1]))

output_file.write(out[1:-1])