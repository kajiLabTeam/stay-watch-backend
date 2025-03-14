import numpy as np
from app.api.schemas import ClusteringResult
from sklearn.mixture import GaussianMixture


def clustering(data: list[int]) -> list[ClusteringResult]:
    d = np.array(data)
    # BIC(ベイズ情報量基準)を用いて最適なクラスタ数を探索
    bic_values = []
    max_clusters = 4
    for n_clusters in range(1, max_clusters + 1):
        gmm = GaussianMixture(n_components=n_clusters, init_params="k-means++")
        gmm.fit(d.reshape(-1, 1))
        bic_values.append(gmm.bic(d.reshape(-1, 1)))
    optimal_componets = np.argmin(bic_values) + 1
    # クラスタリング
    gmm = GaussianMixture(n_components=optimal_componets, init_params="k-means++")
    gmm.fit(d.reshape(-1, 1))
    clusters = make_results_list(d, gmm)
    return clusters


def make_results_list(
    data: np.ndarray, gmm_results: GaussianMixture
) -> list[ClusteringResult]:
    labels = gmm_results.predict(data.reshape(-1, 1))
    centroids = np.array(gmm_results.means_)
    n_clusters = len(centroids)
    # クラスタごとにデータを分割
    cluster: list[list[float]] = [[] for _ in range(n_clusters)]
    for i, label in enumerate(labels):
        cluster[label].append(data[i])
    clusters: list[ClusteringResult] = []
    for i in range(len(cluster)):
        clusters.append(
            ClusteringResult(data=cluster[i], center=np.ravel(centroids[i])[0])
        )
    return clusters
