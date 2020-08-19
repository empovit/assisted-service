import argparse
import os

import deploy_tls_secret
import deployment_options
import utils

SERVICE = "assisted-service"


def main():
    parser = argparse.ArgumentParser()
    deploy_options = deployment_options.load_deployment_options(parser)

    utils.set_profile(deploy_options.target, deploy_options.profile)

    src_file = os.path.join(os.getcwd(), "deploy/assisted-service-service.yaml")
    dst_file = os.path.join(os.getcwd(), "build/assisted-service-service.yaml")
    with open(src_file, "r") as src:
        with open(dst_file, "w+") as dst:
            data = src.read()
            data = data.replace('REPLACE_NAMESPACE', deploy_options.namespace)
            print("Deploying {}".format(dst_file))
            dst.write(data)

    utils.apply(dst_file)

    # in case of OpenShift deploy ingress as well
    if deploy_options.target == utils.INGRESS_REMOTE_TARGET:

        hostname = utils.get_service_host(
            'assisted-installer',
            deploy_options.target,
            deploy_options.domain,
            deploy_options.namespace,
            deploy_options.profile
        )

        if deploy_options.disable_tls:
            template = "assisted-installer-ingress.yaml"
        else:
            print("WARNING: On OpenShift, in order to change TLS redirection behavior update "
                  "spec/tls/insecureEdgeTerminationPolicy (None|Allow|Redirect) "
                  "in the corresponding OpenShift route")
            deploy_tls_secret.generate_secret(output_dir=os.path.join(os.getcwd(), "build"),
                                              service=SERVICE, san=f'DNS:{hostname}',
                                              namespace=deploy_options.namespace)
            template = "assisted-installer-ingress-tls.yaml"

        deploy_ingress(hostname=hostname, namespace=deploy_options.namespace, template_file=template)

    elif not deploy_options.disable_tls:

        hostname = utils.get_service_host(
            SERVICE,
            deploy_options.target,
            deploy_options.domain,
            deploy_options.namespace,
            deploy_options.profile
        )

        deploy_tls_secret.generate_secret(output_dir=os.path.join(os.getcwd(), "build"),
                                          service=SERVICE, san=f'IP:{hostname}',
                                          namespace=deploy_options.namespace, keep_files=True)


def deploy_ingress(hostname, namespace, template_file):
    src_file = os.path.join(os.getcwd(), "deploy", template_file)
    dst_file = os.path.join(os.getcwd(), "build", template_file)
    with open(src_file, "r") as src:
        with open(dst_file, "w+") as dst:
            data = src.read()
            data = data.replace('REPLACE_NAMESPACE', namespace)
            data = data.replace("REPLACE_HOSTNAME", hostname)
            print("Deploying {}".format(dst_file))
            dst.write(data)
    utils.apply(dst_file)


if __name__ == "__main__":
    main()
