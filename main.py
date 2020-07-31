import os
import sys

from plumbum import cli


class F10_CLI(cli.Application):
    PROGNAME = "F10-CLI"
    VERSION = "0.1"
    branch = "master"
    # create = cli.Flag(['create', "c"], help="创建版本描述文件")
    update = cli.Flag(['update', "u"], help="更新版本描述文件")

    @cli.switch(['create', "c"], str)
    def create_version_desc_file(self, task_no):
        self.set_version_file_path(task_no)

        print("create version desc file ")
        print("taks_no : %s" % task_no)
        #  判断版本描述文件是否存在
        if self._check_version_desc_exist(task_no):
            print("version file already exist !")
            return False

        with open(self.version_file_path, 'w', encoding='UTF-8') as f, open(
                sys.path[0] + "/template/version.md", encoding='UTF-8') as t:
            template = t.read()
            # replace template var
            template = self.replace_template(task_no, template)
            f.write(template)
            print("F10 CLI TEST")
        return True

    @cli.switch(['branch'], str)
    def create_version_desc_file(self, branch):
        self.branch = branch

    def replace_template(self, task_no, template):
        version_info = self.get_version_info()
        template = template.replace("{{ TASK_NO }}", task_no)
        template = template.replace("{{ DATE }}", "2020-07-31")
        template = template.replace("{{ DATETIME }}", "2020-07-31 15:27:31")
        template = template.replace("{{ DEVELOPMENT }}", version_info['author'])
        template = template.replace("{{ PROJECT_GIT_NAME }}", "")
        template = template.replace("{{ PROJECT_GIT_URL }}", "")
        return template

    def set_version_file_path(self, task_no):
        self.version_file_path = "./version/" + task_no + ".md"
        return self.version_file_path

    def get_git_diff_log(self):
        commod = 'git log ...' + self.branch + ' --format="%H -|---|- %ci -|---|- %ce -|---|- %s"'
        dir = os.getegid()

    @cli.switch(['check'], str)
    def check_version_desc_exist(self, task_no):
        # 验证文件是否存在

        # 验证文件是否已发布过

        # 已经发布过校验是否有重发参数

        return True

    def _check_version_desc_exist(self, task_no):
        filename = 'version.md'
        if os.path.exists(self.version_file_path):
            return True

        if not os.path.exists('./version/'):
            os.mkdir('./version/')
        return False

    def get_work_dir(self):
        print("获取工作路径")

    # 生成版本描述文件
    def main(self, *args):
        """
        :param args:
        :return:
        """
        print("CLI RUN END")

    def get_version_info(self):
        """
        获取提交版本数据
        :return:
        """

        pass


if __name__ == '__main__':
    print(__file__)
    print(sys.path[0])
    F10_CLI.run()
    # print(__file__)
    # print(sys.argv)
