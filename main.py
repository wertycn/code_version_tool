from plumbum import cli, colors


class F10Cli(cli.Application):
    PROGNAME = colors.green |  "F10-CLI"
    VERSION  =  colors.green | "v1.0.0"
    COLOR_GROUPS = {"Meta-switches": colors.bold & colors.yellow}
    opts = cli.Flag("--ops", help=colors.magenta | "This is help")
    verbose = cli.Flag(["v", "verbose"], help="If given, I will be very talkative")
    update = cli.Flag(['-u','-update','--u','--update'], help="If given, I will be very talkative")

    def main(self, args):
        print("I will n w read {0}".format(args))
        if self.verbose:
            print("Yadda " * 200)

    @cli.autoswitch()
    def update(self,):
        if self.update:
            print('')
        pass


if __name__ == '__main__':
    print("欢迎使用F10终端工具箱 %s" % F10Cli.VERSION)
    F10Cli.run()
